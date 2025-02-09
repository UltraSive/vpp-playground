package main

import (
	"fmt"
	"log"

	"go.fd.io/govpp/binapi/gre"
	interfaces "go.fd.io/govpp/binapi/interface"
	"go.fd.io/govpp/core"
)

func main() {
	conn, err := core.Connect("")
	if err != nil {
		log.Fatalf("Failed to connect to VPP: %v", err)
	}
	defer conn.Disconnect()

	ch, err := conn.NewAPIChannel()
	if err != nil {
		log.Fatalf("Failed to create API channel: %v", err)
	}
	defer ch.Close()

	fmt.Println("Connected to VPP")

	tunnel := gre.GreTunnelAddDel{
		IsAdd:        1,
		Type:         0,
		Instance:     ^uint32(0),
		OuterFibID:   0,
		OuterTableID: 0,
		Tei:          0,
		TunnelSrc:    [4]byte{192, 168, 1, 1},
		TunnelDst:    [4]byte{192, 168, 177, 3}, // Second VM host-connected network
	}

	reply := &gre.GreTunnelAddDelReply{}
	err = ch.SendRequest(&tunnel).ReceiveReply(reply)
	if err != nil {
		log.Fatalf("Failed to create GRE tunnel: %v", err)
	}

	fmt.Printf("Created GRE Tunnel with SwIfIndex: %d\n", reply.SwIfIndex)

	swIfIndex := reply.SwIfIndex
	upReq := interfaces.SwInterfaceSetFlags{
		SwIfIndex: swIfIndex,
		AdminUp:   1,
	}
	upReply := &interfaces.SwInterfaceSetFlagsReply{}

	err = ch.SendRequest(&upReq).ReceiveReply(upReply)
	if err != nil {
		log.Fatalf("Failed to set interface up: %v", err)
	}

	fmt.Println("GRE tunnel is up")
}
