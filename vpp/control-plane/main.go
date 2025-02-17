// Copyright (c) 2017 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// simple-client is an example VPP management application that exercises the
// govpp API on real-world use-cases.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"go.fd.io/govpp"
	"go.fd.io/govpp/adapter/socketclient"
	"go.fd.io/govpp/api"
	interfaces "go.fd.io/govpp/binapi/interface"
	"go.fd.io/govpp/binapi/interface_types"
	"go.fd.io/govpp/binapi/ip"
	"go.fd.io/govpp/binapi/ip_types"
	"go.fd.io/govpp/binapi/vpe"
	"go.fd.io/govpp/core"
	"go.fd.io/govpp/binapi/gre"
	"go.fd.io/govpp/binapi/fib_types"
)

var (
	sockAddr = flag.String("sock", socketclient.DefaultSocketName, "Path to VPP binary API socket file")
)

func main() {
	flag.Parse()

	fmt.Println("Starting simple client example")

	// connect to VPP
	conn, connEv, err := govpp.AsyncConnect(*sockAddr, core.DefaultMaxReconnectAttempts, core.DefaultReconnectInterval)
	if err != nil {
		log.Fatalln("ERROR:", err)
	}
	defer conn.Disconnect()

	e := <-connEv
	if e.State != core.Connected {
		log.Fatalln("ERROR: connecting to VPP failed:", e.Error)
	}

	// check message compatibility
	ch, err := conn.NewAPIChannel()
	if err != nil {
		log.Fatalln("ERROR: creating channel failed:", err)
	}
	defer ch.Close()

	if err := ch.CheckCompatiblity(vpe.AllMessages()...); err != nil {
		log.Fatalf("compatibility check failed: %v", err)
	}
	if err := ch.CheckCompatiblity(interfaces.AllMessages()...); err != nil {
		log.Printf("compatibility check failed: %v", err)
	}

	// process errors encountered during the example
	defer func() {
		if len(errors) > 0 {
			log.Fatalf("finished with %d errors", len(errors))
		}
	}()

	// use Channel request/reply (channel API)
	getVppVersion(ch)
	getSystemTime(ch)
	//idx := createLoopback(ch)
	var idx interface_types.InterfaceIndex = 1 // interface index with correct type
	listInterfaces(ch)
	// Set MAC address for interface
	if err := setInterfaceMAC(ch, idx, "bc:24:11:2c:bd:9a"); err != nil {
		logError(err, "setting MAC address")
	}
	addIPAddress(ch, idx)
	listIPaddresses(ch, idx)
	//watchInterfaceEvents(ch, idx)
	addIPRoute(ch, "0.0.0.0/0", "172.93.110.1")
	setInterfaceStatus(ch, idx, true)

	// Set MTU for interface
	if err := setInterfaceMTU(ch, idx, 1500); err != nil {
		logError(err, "setting MTU")
	}

	// Create GRE tunnel
	tunnelIdx, err := createGRETunnel(ch, "172.93.110.120", "45.76.233.197")
	if err != nil {
		logError(err, "creating GRE tunnel")
		return
	}

	// Add IP address to the tunnel interface
	if err := addInterfaceIPAddress(ch, tunnelIdx, "10.10.10.1/30"); err != nil {
		logError(err, "adding IP address to tunnel")
		return
	}

	// Bring the tunnel interface up
	setInterfaceStatus(ch, tunnelIdx, true)

	// Add route through the GRE tunnel
	if err := addIPRouteViaInterface(ch, "103.195.102.92/24", "10.10.10.2", tunnelIdx); err != nil {
		logError(err, "adding route via GRE tunnel")
	}
}

func getVppVersion(ch api.Channel) {
	req := &vpe.ShowVersion{}
	reply := &vpe.ShowVersionReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		logError(err, "retrieving version")
		return
	}

	fmt.Printf("VPP version: %q\n", reply.Version)
}

func getSystemTime(ch api.Channel) {
	req := &vpe.ShowVpeSystemTime{}
	reply := &vpe.ShowVpeSystemTimeReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		logError(err, "retrieving system time")
		return
	}

	fmt.Printf("system time: %v\n", reply.VpeSystemTime)
}

func createLoopback(ch api.Channel) interface_types.InterfaceIndex {
	req := &interfaces.CreateLoopback{}
	reply := &interfaces.CreateLoopbackReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		logError(err, "creating loopback")
		return 0
	}

	fmt.Printf("loopback created: %v\n", reply.SwIfIndex)

	return reply.SwIfIndex
}

func listInterfaces(ch api.Channel) {
	reqCtx := ch.SendMultiRequest(&interfaces.SwInterfaceDump{
		SwIfIndex: ^interface_types.InterfaceIndex(0),
	})
	for {
		iface := &interfaces.SwInterfaceDetails{}
		stop, err := reqCtx.ReceiveReply(iface)
		if stop {
			break
		}
		if err != nil {
			logError(err, "listing interfaces")
			return
		}
		fmt.Printf(" - interface: %+v (ifIndex: %v)\n", iface.InterfaceName, iface.SwIfIndex)
		marshal(iface)
	}

	fmt.Println("OK")
	fmt.Println()
}

func addIPAddress(ch api.Channel, ifIdx interface_types.InterfaceIndex) {
	addr := ip_types.NewAddress(net.IPv4(172, 93, 110, 120))

	req := &interfaces.SwInterfaceAddDelAddress{
		SwIfIndex: ifIdx,
		IsAdd:     true,
		Prefix:    ip_types.AddressWithPrefix{Address: addr, Len: 24},
	}
	marshal(req)
	reply := &interfaces.SwInterfaceAddDelAddressReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		logError(err, "adding IP address")
		return
	}
}

func listIPaddresses(ch api.Channel, index interface_types.InterfaceIndex) {
	reqCtx := ch.SendMultiRequest(&ip.IPAddressDump{
		SwIfIndex: index,
	})
	for {
		ipAddr := &ip.IPAddressDetails{}
		stop, err := reqCtx.ReceiveReply(ipAddr)
		if err != nil {
			logError(err, "listing IP addresses")
			return
		}
		if stop {
			break
		}
		fmt.Printf(" - IP address: %+v\n", ipAddr)
		marshal(ipAddr)
	}
}

func setInterfaceStatus(ch api.Channel, ifIdx interface_types.InterfaceIndex, up bool) {
	var flags interface_types.IfStatusFlags
	if up {
		flags = interface_types.IF_STATUS_API_FLAG_ADMIN_UP
	} else {
		flags = 0
	}
	if err := ch.SendRequest(&interfaces.SwInterfaceSetFlags{
		SwIfIndex: ifIdx,
		Flags:     flags,
	}).ReceiveReply(&interfaces.SwInterfaceSetFlagsReply{}); err != nil {
		log.Fatalln("ERROR:  setting interface flags failed:", err)
	}
	if up {
		fmt.Printf("interface status set to UP")
	} else {
		fmt.Printf("interface status set to DOWN")
	}
}

func marshal(v interface{}) {
	fmt.Printf("GO: %#v\n", v)
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("JSON: %s\n", b)
}

var errors []error

func logError(err error, msg string) {
	fmt.Printf("ERROR: %s: %v\n", msg, err)
	errors = append(errors, err)
}

func addIPRoute(ch api.Channel, dstAddr string, nextHopAddr string) {
	// Parse the destination CIDR
	_, dst, err := net.ParseCIDR(dstAddr)
	if err != nil {
		logError(err, "parsing destination CIDR")
		return
	}

	// Parse the next hop address
	nextHop := net.ParseIP(nextHopAddr)
	if nextHop == nil {
		logError(fmt.Errorf("invalid next hop address"), "parsing next hop address")
		return
	}

	ones, _ := dst.Mask.Size()

	// Create route path using fib_types.FibPath
	path := fib_types.FibPath{
		SwIfIndex: ^uint32(0), // Use default interface index
		Proto:     fib_types.FibPathNhProto(ip_types.ADDRESS_IP4),
		Weight:    1,
		Nh: fib_types.FibPathNh{
			Address: ip_types.AddressUnion{
				XXX_UnionData: toAddress(nextHop).Un.XXX_UnionData,
			},
		},
	}

	// Prepare route add request
	req := &ip.IPRouteAddDel{
		IsAdd: true,
		Route: ip.IPRoute{
			Prefix: ip_types.Prefix{
				Address: toAddress(dst.IP),
				Len:     uint8(ones),
			},
			NPaths: 1,
			Paths:  []fib_types.FibPath{path},
		},
	}

	// Send the request
	reply := &ip.IPRouteAddDelReply{}
	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		logError(err, "adding IP route")
		return
	}

	fmt.Printf("Route added successfully: %s via %s\n", dstAddr, nextHopAddr)
}

func createGRETunnel(ch api.Channel, srcAddr, dstAddr string) (interface_types.InterfaceIndex, error) {
	// Parse source and destination IP addresses
	src := net.ParseIP(srcAddr)
	if src == nil {
		return 0, fmt.Errorf("invalid source address: %s", srcAddr)
	}
	dst := net.ParseIP(dstAddr)
	if dst == nil {
		return 0, fmt.Errorf("invalid destination address: %s", dstAddr)
	}

	// Create GRE tunnel request
	req := &gre.GreTunnelAddDel{
		IsAdd: true,
		Tunnel: gre.GreTunnel{
			Src:      toAddress(src),
			Dst:      toAddress(dst),
			Instance: ^uint32(0), // Default instance
		},
	}

	reply := &gre.GreTunnelAddDelReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		return 0, fmt.Errorf("failed to create GRE tunnel: %v", err)
	}

	fmt.Printf("GRE tunnel created successfully: src %s dst %s (SwIfIndex: %d)\n", 
		srcAddr, dstAddr, reply.SwIfIndex)

	return reply.SwIfIndex, nil
}

func setInterfaceMAC(ch api.Channel, ifIdx interface_types.InterfaceIndex, macAddr string) error {
	// Parse MAC address
	hwAddr, err := net.ParseMAC(macAddr)
	if err != nil {
		return fmt.Errorf("invalid MAC address %s: %v", macAddr, err)
	}

	// Convert MAC address to VPP format
	var vppHwAddr [6]uint8
	copy(vppHwAddr[:], hwAddr)

	// Create request to set MAC address
	req := &interfaces.SwInterfaceSetMacAddress{
		SwIfIndex:  ifIdx,
		MacAddress: vppHwAddr,
	}

	reply := &interfaces.SwInterfaceSetMacAddressReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		return fmt.Errorf("failed to set MAC address: %v", err)
	}

	fmt.Printf("MAC address set successfully: %s for interface %d\n", macAddr, ifIdx)
	return nil
}

func setInterfaceMTU(ch api.Channel, ifIdx interface_types.InterfaceIndex, mtu uint32) error {
	req := &interfaces.SwInterfaceSetMtu{
		SwIfIndex: ifIdx,
		Mtu:       []uint32{mtu}, // MTU as slice for newer VPP versions
	}

	reply := &interfaces.SwInterfaceSetMtuReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		return fmt.Errorf("failed to set MTU: %v", err)
	}

	fmt.Printf("MTU set successfully: %d for interface %d\n", mtu, ifIdx)
	return nil
}

func addIPRouteViaInterface(ch api.Channel, dstAddr string, nextHopAddr string, ifIdx interface_types.InterfaceIndex) error {
	// Parse the destination CIDR
	_, dst, err := net.ParseCIDR(dstAddr)
	if err != nil {
		return fmt.Errorf("parsing destination CIDR: %v", err)
	}

	// Parse the next hop address
	nextHop := net.ParseIP(nextHopAddr)
	if nextHop == nil {
		return fmt.Errorf("invalid next hop address: %s", nextHopAddr)
	}

	ones, _ := dst.Mask.Size()

	// Create route path using fib_types.FibPath
	path := fib_types.FibPath{
		SwIfIndex: uint32(ifIdx),
		Proto:     fib_types.FibPathNhProto(ip_types.ADDRESS_IP4),
		Weight:    1,
		Nh: fib_types.FibPathNh{
			Address: ip_types.AddressUnion{
				XXX_UnionData: toAddress(nextHop).Un.XXX_UnionData,
			},
		},
	}

	// Prepare route add request
	req := &ip.IPRouteAddDel{
		IsAdd: true,
		Route: ip.IPRoute{
			Prefix: ip_types.Prefix{
				Address: toAddress(dst.IP),
				Len:     uint8(ones),
			},
			NPaths: 1,
			Paths:  []fib_types.FibPath{path},
		},
	}

	// Send the request
	reply := &ip.IPRouteAddDelReply{}
	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		return fmt.Errorf("failed to add IP route: %v", err)
	}

	fmt.Printf("Route added successfully: %s via %s through interface %d\n", 
		dstAddr, nextHopAddr, ifIdx)
	return nil
}

func addInterfaceIPAddress(ch api.Channel, ifIdx interface_types.InterfaceIndex, ipAddr string) error {
	// Parse the IP address with subnet
	ip, ipNet, err := net.ParseCIDR(ipAddr)
	if err != nil {
		return fmt.Errorf("invalid IP address or subnet: %v", err)
	}

	// Calculate the prefix length
	ones, _ := ipNet.Mask.Size()

	req := &interfaces.SwInterfaceAddDelAddress{
		SwIfIndex: ifIdx,
		IsAdd:     true,
		Prefix: ip_types.AddressWithPrefix{
			Address: toAddress(ip),
			Len:     uint8(ones),
		},
	}

	reply := &interfaces.SwInterfaceAddDelAddressReply{}

	if err := ch.SendRequest(req).ReceiveReply(reply); err != nil {
		return fmt.Errorf("failed to add IP address: %v", err)
	}

	fmt.Printf("IP address %s added to interface %d\n", ipAddr, ifIdx)
	return nil
}

func toAddress(ip net.IP) ip_types.Address {
	var addr ip_types.Address
	if ip4 := ip.To4(); ip4 != nil {
		var ip4Addr [4]uint8
		copy(ip4Addr[:], ip4)
		addr.Af = ip_types.ADDRESS_IP4
		addr.Un.XXX_UnionData = ip_types.AddressUnion{
			XXX_UnionData: [16]byte{ip4Addr[0], ip4Addr[1], ip4Addr[2], ip4Addr[3]},
		}.XXX_UnionData
	} else {
		var ip6Addr [16]uint8
		copy(ip6Addr[:], ip)
		addr.Af = ip_types.ADDRESS_IP6
		addr.Un.XXX_UnionData = ip_types.AddressUnion{
			XXX_UnionData: [16]byte(ip6Addr),
		}.XXX_UnionData
	}
	return addr
}