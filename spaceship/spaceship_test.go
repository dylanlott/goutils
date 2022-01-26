package spaceship

import (
	"testing"

	"github.com/stretchr/testify"
)

// Testing is incredibly important in any language.
// Test functions in go always follow the form of `TestXXXXX` with the name following.
// To run these tests in the terminal, you would run `go test -v ./spaceship/

func TestSpaceship(t *testing.T) {
}

func Testnew(t *testing.T) {
	ship := NewShip()
	testify.Assert()
}
