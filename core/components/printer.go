// A component that prints its input on screen
package components

import (
	"fmt"

	flow "github.com/trustmaster/goflow"
)

type Printer struct {
	flow.Component
	Line <-chan string // inport
}

// Prints a line when it gets it
func (p *Printer) OnLine(line string) {
	fmt.Println(line)
}
