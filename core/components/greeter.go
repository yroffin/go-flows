// A component that generates greetings
package components

import (
	"fmt"

	flow "github.com/trustmaster/goflow"
)

type Greeter struct {
	flow.Component               // component "superclass" embedded
	Name           <-chan string // input port
	Res            chan<- string // output port
}

// Reaction to a new name input
func (g *Greeter) OnName(name string) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	// send the greeting to the output port
	g.Res <- greeting
}
