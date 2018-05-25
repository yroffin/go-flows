package main

import (
	"fmt"
	"reflect"

	"github.com/trustmaster/goflow"
	"github.com/yroffin/go-flows/core/components"
)

// Our greeting network
type GreetingApp struct {
	flow.Graph // graph "superclass" embedded
}

// Graph constructor and structure definition
func Factory() interface{} {
	instance := new(components.Greeter)
	fmt.Println("Type", reflect.TypeOf(instance))
	reflect.New()
	return instance
}

// Graph constructor and structure definition
func NewGreetingApp() *GreetingApp {
	n := new(GreetingApp) // creates the object in heap
	n.InitGraphState()    // allocates memory for the graph
	// Add processes to the network
	n.Add(Factory(), "greeter")
	n.Add(new(components.Printer), "printer")
	// Connect them with a channel
	n.Connect("greeter", "Res", "printer", "Line")
	// Our net has 1 inport mapped to greeter.Name
	n.MapInPort("In", "greeter", "Name")
	return n
}

func main() {
	// Create the network
	net := NewGreetingApp()
	// We need a channel to talk to it
	in := make(chan string)
	net.SetInPort("In", in)
	// Run the net
	flow.RunNet(net)
	// Now we can send some names and see what happens
	in <- "John"
	in <- "Boris"
	in <- "Hanna"
	// Close the input to shut the network down
	close(in)
	// Wait until the app has done its job
	<-net.Wait()
}
