package main

import (
	"syscall/js"
)

func main() {
	c := make(chan bool)
	//1. Adding an <h1> element in the HTML document
	document := js.Global().Get("document")
	p := document.Call("createElement", "h1")
	p.Set("innerHTML", "Hello from Golang!")
	document.Get("body").Call("appendChild", p)

	//2. Exposing go functions/values in javascript variables.
	js.Global().Set("goVar", "I am a variable set from Go")
	js.Global().Set("sayHello", js.FuncOf(sayHello))

	//3. This channel will prevent the go program to exit
	<-c
}

func sayHello(this js.Value, inputs []js.Value) interface{} {
	firstArg := inputs[0].String()
	return "Hi " + firstArg + " from Go!"
}
