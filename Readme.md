# Go Snoop

**This project is now merged into the google/gopacket project. Updates will be done there**

A libary for read snoop file format v2 with gopacket integration.
See readsnoop for example usage.

## Header

Here some Header informations of this format. For details see rfc1761.
I like to store my notices sry ;-)

```txt

Snoop Header (16 byte)
    8 byte magic (0x706f6f6e73000000)
    4 byte version
    4 byte link type

Data Format: All integer values are stored in big-endian order (high-order first)
packet data:

+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Original Length (uint32)               | size captured packet on network
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Included Length (uint32                | size Data field (Packet Data)
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                      Packet Record Length (uint32)            | size total length rec (24 octets of descriptive information + packet data + pad field)
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                        Cumulative Drops (uint32)              | num of packetes that were lost by the system (lack impl. set to 0)
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Timestamp Seconds (uint32)              | timestamp since 01 01 1970
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                     Timestamp Microseconds (uint32)           | microsecond packet arrival time
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                                                               | variable length holding the packet that was captured, beginning with its datalink header
.                                                               .
.                          Packet Data                          .
.                                                               .
+                                               +- - - - - - - -+
|                                               |     Pad       | variable  length field holding zeros
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
```

## Futher Reading

[Custom Decoder for Gopacket](https://damianzaremba.co.uk/2017/12/decoding-arista-ethertype-headers-with-gopacket/)

[Basic Usage of gopacket and pcapgo](https://gowalker.org/github.com/google/gopacket/pcapgo)

[Devdungeon Gopacket](https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket)