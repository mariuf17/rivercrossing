package state

var state = "[ulv sau kålhode HS ---\\ \\--/---------------/---]"

func EnterBoat() {
	state = "[ulv sau kålhode ---\\ \\_HS_/-----------------/---]"
}
func Rivercross() {
	state = "[ulv sau kålhode ---\\ \\------------------\\-HS-/---]"
}
func ViewState() string {
	return state
}
