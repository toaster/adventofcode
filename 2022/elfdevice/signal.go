package elfdevice

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

// Packet represents a packet of a signal
type Packet []packetItem

// ParsePacket parses a packet from a line of input
func ParsePacket(input string) Packet {
	packet, _ := parsePacket(input)
	return packet
}

// Compare returns the order of the Packet compared to another one.
// That is -1 if the Packet is less than, 0 if it is equal to, and 1 if it is greater than the other one.
func (p Packet) Compare(other Packet) int {
	return p.compare(other, "")
}

func (p Packet) compare(other Packet, indent string) int {
	// fmt.Printf("%s- Compare %s vs %s\n", indent, p.string(), other.string())
	indent += "  "
	for i, item := range p {
		if i >= len(other) {
			fmt.Printf("%s- Right side ran out of items, so inputs are not in the right order (1)\n", indent)
			return 1
		}

		itemResult := item.compare(other[i], indent)
		if itemResult != 0 {
			return itemResult
		}
	}
	if len(other) == len(p) {
		return 0
	}

	fmt.Printf("%s- Left side ran out of items, so inputs are in the right order (-1)\n", indent)
	return -1
}

func (p Packet) string() string {
	s := strings.Builder{}
	s.WriteRune('[')
	for i, item := range p {
		if i > 0 {
			s.WriteRune(',')
		}
		if item.packet != nil {
			s.WriteString(item.packet.string())
		} else {
			s.WriteString(strconv.Itoa(item.value))
		}
	}
	s.WriteRune(']')
	return s.String()
}

func parsePacket(input string) (packet Packet, remainder string) {
	remainder = input[1:]
	if remainder[0] == ']' {
		packet = Packet{}
		remainder = remainder[1:]
		return
	}

	done := false
	for !done {
		var item packetItem
		item, remainder = parseItem(remainder)
		packet = append(packet, item)
		done = remainder[0] == ']'
		remainder = remainder[1:]
	}
	return
}

func parseItem(input string) (packetItem, string) {
	if input[0] == '[' {
		packet, remainder := parsePacket(input)
		return packetItem{packet: packet}, remainder
	}

	value, remainder := parseValue(input)
	return packetItem{value: value}, remainder
}

func parseValue(input string) (value int, remainder string) {
	for i, c := range input {
		if c == ',' || c == ']' {
			remainder = input[i:]
			input = input[:i]
			break
		}
	}
	value = io.ParseInt(input)
	return
}

type packetItem struct {
	packet Packet
	value  int
}

func (i packetItem) compare(other packetItem, indent string) int {
	if i.packet == nil && other.packet == nil {
		fmt.Printf("%s- Compare %d vs %d\n", indent, i.value, other.value)
		if i.value < other.value {
			fmt.Printf("%s  - Left side is smaller, so inputs are in the right order (-1)\n", indent)
			return -1
		}
		if i.value == other.value {
			return 0
		}
		fmt.Printf("%s  - Right side is smaller, so inputs are not in the right order (1)\n", indent)
		return 1
	}

	packet := i.packet
	if packet == nil {
		packet = Packet{i}
	}
	otherPacket := other.packet
	if otherPacket == nil {
		otherPacket = Packet{other}
	}
	return packet.compare(otherPacket, indent)
}
