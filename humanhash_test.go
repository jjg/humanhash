package humanhash

import (
	"testing"
)

func TestHumanize(t *testing.T) {
	want := "12345678910"
	if want != Humanize(want) {
		t.Fatal("Dman!")
	}
}
