package main

import (
	"fmt"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	var binary []uint8
	input := io.ReadLines()[0]
	fmt.Println("input:", input)
	for i := 0; i < len(input); i += 2 {
		binary = appendHex(binary, input[i:i+2])
	}
	fmt.Print("binary: ")
	for _, n := range binary {
		fmt.Printf("%08b", n)
	}
	fmt.Println()

	p, _ := parsePacket(binary, 0)
	printPacket(p, "")
	fmt.Println(p.value)
}

func printPacket(p *packet, prefix string) {
	fmt.Printf("%sVersion: %d, Type: %d, Value: %d", prefix, p.version, p.typeID, p.value)
	if p.subPackets != nil {
		fmt.Printf(", Subpackets Version Sum: %d, Subpackets: [\n", p.subPacketsVersionSum)
		for _, subPacket := range p.subPackets {
			printPacket(subPacket, prefix+"  ")
		}
		fmt.Print(prefix + "]")
	}
	fmt.Println()
}

type packet struct {
	version              uint8
	typeID               uint8
	value                uint64
	subPacketsVersionSum uint64
	subPackets           []*packet
}

func parsePacket(binary []uint8, offset int) (*packet, int) {
	p := &packet{}
	p.version, offset = readByte(binary, offset, 3)
	p.typeID, offset = readByte(binary, offset, 3)
	switch p.typeID {
	case 4:
		p.value, offset = readValue(binary, offset)
	default:
		var lengthType uint8
		lengthType, offset = readByte(binary, offset, 1)
		if lengthType == 1 {
			var packetCount uint16
			packetCount, offset = readWord(binary, offset, 11)
			for i := uint16(0); i < packetCount; i++ {
				offset = parseSubpacket(binary, offset, p)
			}
		} else {
			var packetLength uint16
			packetLength, offset = readWord(binary, offset, 15)
			maxOffset := offset + int(packetLength)
			for offset < maxOffset {
				offset = parseSubpacket(binary, offset, p)
			}
		}
		switch p.typeID {
		case 0:
			for _, subPacket := range p.subPackets {
				p.value += subPacket.value
			}
		case 1:
			p.value = 1
			for _, subPacket := range p.subPackets {
				p.value *= subPacket.value
			}
		case 2:
			for _, subPacket := range p.subPackets {
				if p.value == 0 || subPacket.value < p.value {
					p.value = subPacket.value
				}
			}
		case 3:
			for _, subPacket := range p.subPackets {
				if subPacket.value > p.value {
					p.value = subPacket.value
				}
			}
		case 5:
			if p.subPackets[0].value > p.subPackets[1].value {
				p.value = 1
			}
		case 6:
			if p.subPackets[0].value < p.subPackets[1].value {
				p.value = 1
			}
		case 7:
			if p.subPackets[0].value == p.subPackets[1].value {
				p.value = 1
			}
		}
	}
	return p, offset
}

func parseSubpacket(binary []uint8, offset int, p *packet) int {
	var sp *packet
	sp, offset = parsePacket(binary, offset)
	p.subPackets = append(p.subPackets, sp)
	p.subPacketsVersionSum += uint64(sp.version) + sp.subPacketsVersionSum
	return offset
}

func readWord(binary []uint8, offset int, bits int) (uint16, int) {
	var b uint8
	b, offset = readByte(binary, offset, bits-8)
	var packetLength uint16
	packetLength = uint16(b) << 8
	b, offset = readByte(binary, offset, 8)
	packetLength |= uint16(b)
	return packetLength, offset
}

func readValue(binary []uint8, offset int) (uint64, int) {
	var value uint64
	for {
		var indicator uint8
		indicator, offset = readByte(binary, offset, 1)
		var nibble uint8
		nibble, offset = readByte(binary, offset, 4)
		value <<= 4
		value |= uint64(nibble)
		if indicator == 0 {
			break
		}
	}
	return value, offset
}

func readByte(binary []uint8, offset int, bitCount int) (uint8, int) {
	byteOffset := offset / 8
	bitOffset := offset % 8
	var value uint8
	usableBits := 8 - bitOffset
	if usableBits < bitCount {
		neededFromNextByte := bitCount - usableBits
		value = binary[byteOffset] << neededFromNextByte
		value |= binary[byteOffset+1] >> (8 - neededFromNextByte)
	} else {
		value = binary[byteOffset] >> (usableBits - bitCount)
	}
	var mask uint8
	for i := 0; i < bitCount; i++ {
		mask <<= 1
		mask |= 1
	}
	return value & mask, offset + bitCount
}

func appendHex(binary []uint8, hexByte string) []uint8 {
	if len(hexByte) == 1 {
		hexByte = hexByte + "0"
	}
	decodedByte := decodeNibble(hexByte[0])
	decodedByte <<= 4
	decodedByte += decodeNibble(hexByte[1])
	return append(binary, decodedByte)
}

func decodeNibble(b uint8) uint8 {
	switch b {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'A', 'a':
		return 10
	case 'B', 'b':
		return 11
	case 'C', 'c':
		return 12
	case 'D', 'd':
		return 13
	case 'E', 'e':
		return 14
	case 'F', 'f':
		return 15
	}
	return 0
}
