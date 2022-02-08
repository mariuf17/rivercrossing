package main

import (
	"fmt"
	"github.com/mariuf17/rivercrossing/State"
)

func main() {
	fmt.Println(state.ViewState())

	state.EnterBoat()
	fmt.Println(state.ViewState())

	state.Rivercross()
	fmt.Println(state.ViewState())

}
