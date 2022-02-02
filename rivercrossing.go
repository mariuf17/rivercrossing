package main

import (
	"fmt"
	state "rivercrossing"
)

func main() {

}

func main() {
	fmt.Println(state.ViewState())

	state.EnterBoat()
	fmt.Println(state.ViewState())

	state.Rivercross()
	fmt.Println(state.ViewState())

}
