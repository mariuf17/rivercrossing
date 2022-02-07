package state

import (
	"testing"
)

func TestCheckState(t *testing.T) {
	wanted := "[ulv sau kålhode HS ---\\ \\--/---------------/---]"
	state := ViewState()
	if state != wanted {
		t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
