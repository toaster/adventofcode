package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

const (
	low  pulseType = false
	high pulseType = true
)

func main() {
	queue := make(chan *pulse, 1000000)
	modules := map[string]communicator{}
	for _, line := range io.ReadLines() {
		idAndDests := strings.Split(line, " -> ")
		id := idAndDests[0]
		dests := strings.Split(idAndDests[1], ", ")
		switch id[0] {
		case 'b':
			modules[id] = &broadcast{module: module{
				communicators: modules,
				destinations:  dests,
				n:             id,
				queue:         queue,
			}}
		case '%':
			name := id[1:]
			modules[name] = &flipFlop{module: module{
				communicators: modules,
				destinations:  dests,
				n:             name,
				queue:         queue,
			}}
		case '&':
			name := id[1:]
			modules[name] = &conjunction{
				module: module{
					communicators: modules,
					destinations:  dests,
					n:             name,
					queue:         queue,
				},
				memory: map[communicator]pulseType{},
			}
		}
	}

	for _, c := range modules {
		for _, n := range c.destinationNames() {
			d := modules[n]
			if d == nil {
				modules[n] = &output{
					module: module{n: n},
				}
			} else if con, ok := d.(*conjunction); ok {
				con.memory[c] = low
			}
		}
	}

	buttonPresses := 0
	rxLow := false
	for !rxLow {
		buttonPresses++
		queue <- &pulse{receiver: modules["broadcaster"], typ: low}
		for len(queue) > 0 {
			p := <-queue
			if p.receiver == nil {
				panic("oops")
			}
			if p.receiver != nil {
				if p.receiver.name() == "rx" {
					// fmt.Println(buttonPresses, p.receiver.name(), p.typ)
				}
				if p.receiver.name() == "rx" && p.typ == low {
					rxLow = true
				}
				p.receiver.receive(p)
			}
		}
	}
	fmt.Println(buttonPresses)
}

type broadcast struct {
	module
}

var _ communicator = (*broadcast)(nil)

func (b *broadcast) receive(p *pulse) {
	b.send(b, p.typ)
}

type conjunction struct {
	module
	memory map[communicator]pulseType
}

var _ communicator = (*conjunction)(nil)

func (c *conjunction) receive(p *pulse) {
	c.memory[p.sender] = p.typ
	typ := low
	for _, t := range c.memory {
		if t == low {
			typ = high
			break
		}
	}
	c.send(c, typ)
}

type communicator interface {
	destinationNames() []string
	name() string
	receive(t *pulse)
}

type flipFlop struct {
	module
	isOn bool
}

var _ communicator = (*flipFlop)(nil)

func (f *flipFlop) receive(p *pulse) {
	if p.typ == high {
		return
	}

	var typ pulseType
	if f.isOn {
		f.isOn = false
		typ = low
	} else {
		f.isOn = true
		typ = high
	}
	f.send(f, typ)
}

type module struct {
	communicators map[string]communicator
	destinations  []string
	n             string
	queue         chan *pulse
}

func (m *module) destinationNames() []string {
	return m.destinations
}

func (m *module) name() string {
	return m.n
}

func (m *module) send(sender communicator, typ pulseType) {
	for _, destination := range m.destinations {
		m.queue <- &pulse{receiver: m.communicators[destination], sender: sender, typ: typ}
	}
}

type output struct {
	module
}

var _ communicator = (*output)(nil)

func (o *output) receive(_ *pulse) {
}

type pulse struct {
	receiver communicator
	sender   communicator
	typ      pulseType
}

type pulseType bool

func (t pulseType) String() string {
	if t == low {
		return "low"
	}
	return "high"
}
