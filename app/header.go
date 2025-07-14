package main

import "encoding/binary"

type MessageHeader struct {
	packetID uint16

	flag uint16

	qdCount uint16
	anCount uint16
	nsCount uint16
	arCount uint16
}

/*
*

	//   [15]    QR:   Query/Response (0=Query, 1=Response)
	//   [14-11] OPCODE: Operation type (4 bits)
	//   [10]    AA:   Authoritative Answer (1 bit)
	//   [9]     TC:   Truncation flag (1 bit)
	//   [8]     RD:   Recursion Desired (1 bit)
	//   [7]     RA:   Recursion Available (1 bit)
	//   [6-4]   Z:    Reserved (3 bits, must be 0)
	//   [3-0]   RCODE: Response code (4 bits)
*/
const (
	flagQR     = 0xF
	flagOPCODE = 0xB
	flagAA     = 0xA
	flagTC     = 0x9
	flagRD     = 0x8
	flagRA     = 0x7
	flagZ      = 0x4
	flagRCODE  = 0x0
)

func ParseHeader(data [12]byte) *MessageHeader {
	h := &MessageHeader{}
	h.packetID = binary.BigEndian.Uint16(data[0:2])
	h.flag = binary.BigEndian.Uint16(data[2:4])
	h.qdCount = binary.BigEndian.Uint16(data[4:6])
	h.anCount = binary.BigEndian.Uint16(data[6:8])
	h.nsCount = binary.BigEndian.Uint16(data[8:10])
	h.arCount = binary.BigEndian.Uint16(data[10:12])
	return h
}

func (h *MessageHeader) ToBytes() []byte {
	b := make([]byte, 12)
	binary.BigEndian.AppendUint16(b, h.packetID)
	binary.BigEndian.AppendUint16(b, h.flag)
	binary.BigEndian.AppendUint16(b, h.qdCount)
	binary.BigEndian.AppendUint16(b, h.anCount)
	binary.BigEndian.AppendUint16(b, h.nsCount)
	binary.BigEndian.AppendUint16(b, h.arCount)
	return b
}

func NewDefaultHeader() *MessageHeader {
	return &MessageHeader{
		packetID: 1234,
		flag:     0x80,
		qdCount:  0,
		anCount:  0,
		nsCount:  0,
		arCount:  0,
	}
}
