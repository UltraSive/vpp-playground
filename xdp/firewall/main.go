package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
	"unsafe"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

// Define the Go equivalents of the eBPF structures
type ActionType uint32

const (
	ALLOW ActionType = iota
	BLOCK
	RATE_LIMIT
)

type DestUsage uint32

const (
	LOCAL_MACHINE DestUsage = iota
	TUNNEL_HOST
	TUNNEL_FORWARD
	FACILITY
)

type DestinationInfo struct {
	Usage         DestUsage
	DefaultAction ActionType
	RxBytes       uint32
	RxPkts        uint32
	TxBytes       uint32
	TxPkts        uint32
}

// Convert IPv4 address in dotted-decimal format to uint32
func ipToUint32(ipStr string) (uint32, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", ipStr)
	}

	// Ensure the IP address is an IPv4 address
	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("not an IPv4 address: %s", ipStr)
	}

	return binary.BigEndian.Uint32(ip), nil
}

func main() {
	// Remove resource limits for kernels <5.11.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock:", err)
	}

	// Load the compiled eBPF ELF and load it into the kernel.
	var objs tunnelObjects
	if err := loadTunnelObjects(&objs, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	ifname := "eth0" // Change this to an interface on your machine.
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		log.Fatalf("Getting interface %s: %s", ifname, err)
	}

	// Attach count_packets to the network interface.
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.XdpSockProg,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatal("Attaching XDP:", err)
	}
	defer link.Close()

	loadDestinationMap(objs)

	log.Printf("Counting incoming packets on %s..", ifname)

	// Periodically fetch the packet counter from PktCount,
	// exit the program when interrupted.
	tick := time.Tick(time.Second)
	stop := make(chan os.Signal, 5)
	signal.Notify(stop, os.Interrupt)
	for {
		select {
		case <-tick:
			var count uint64
			err := objs.PktCount.Lookup(uint32(0), &count)
			if err != nil {
				log.Fatal("Map lookup:", err)
			}
			log.Printf("Received %d packets", count)
		case <-stop:
			log.Print("Received signal, exiting..")
			return
		}
	}
}

func loadDestinationMap(objs tunnelObjects) {
	// Access the map from counterMaps
	destinationMap := objs.Ipv4DestinationMap

	// Example: Write a value to the map
	keyStr := "103.195.102.92"     // Example IP address
	key, err := ipToUint32(keyStr) // Example key
	if err != nil {
		log.Fatalf("failed to convert IP to uint32: %v", err)
	}
	value := DestinationInfo{
		Usage:         TUNNEL_FORWARD,
		DefaultAction: ALLOW,
		RxBytes:       0,
		RxPkts:        0,
		TxBytes:       0,
		TxPkts:        0,
	}

	if err := destinationMap.Update(unsafe.Pointer(&key), unsafe.Pointer(&value), 0); err != nil {
		log.Fatalf("failed to update firewall map: %v", err)
	}

	log.Println("Successfully updated the destination info map")

	// Example: Read a value from the map
	/*var readValue uint32
	if err := destinationMap.Lookup(unsafe.Pointer(&key), unsafe.Pointer(&readValue)); err != nil {
		log.Fatalf("failed to lookup map: %v", err)
	}

	log.Printf("Read value: %d\n", readValue)*/
}
