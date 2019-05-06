package main

import (
	"fmt"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/snoop"
)

func main() {
	//download snoop from https://wiki.wireshark.org/SampleCaptures
	f, err := os.Open("example.snoop")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	handle, err := snoop.NewReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	lt, err := handle.LinkType()
	if err != nil {
		fmt.Println(err)
		return
	}
	packetSource := gopacket.NewPacketSource(handle, lt)

	cnt := 0
	for packet := range packetSource.Packets() {
		fmt.Println(packet.Metadata())
		fmt.Println(packet)
		cnt++
	}
	fmt.Printf("Packet count: %d\n", cnt)
}
