package spaceship

import (
	"testing"

	"github.com/stretchr/testify/assert" // this package just makes it easy for us to assert that things are what we expect them to be
)

// Testing is incredibly important in any language.
// Test functions in go always follow the form of `TestXXXXX` with the name following.
// To run these tests in the terminal, you would run `go test -v ./spaceship/

// TestNew tests the Constructor function of our ship.
func TestNew(t *testing.T) {
	got := NewShip()
	wanted := SimpleChassis{
		Engine:     &FusionEngine{},
		Ammunition: []Bullet{{Size: 10, Distance: 100000}},
	}
	assert.Equal(t, wanted, got, "failed to get ship: wanted: %v - got - %v", wanted, got)
}

// TestEngineRun should pass when you pull this down for the first time.
func TestEngineRun(t *testing.T) {
	// First phase of the test - Assemble, sometimes called setup.
	// we're testing our NewShip function, so we invoke it and assign the result.
	ship := NewShip()
	// Next phase of the test: Act
	// We're testing Run in this function so we call the Run function.
	err := ship.Engine.Run()
	// Last phase: Assert.
	// We want to assert that our engine ran and no error was returned.
	assert.Nil(t, err)
}

// TestSetEngine tests our engine hot swapping capability.
// This test will fail
func TestSetEngine(t *testing.T) {
	ship := NewShip()
	ship.SetEngine(&UpgradedFusionEngine{
		// TODO: This test will fail. Edit it to pass.
	})
	assert.Equal(t, &UpgradedFusionEngine{
		FuelType: "U-256",
	}, ship.Engine, "failed to set engine correctly")
}

// TODO: This test will fail. You'll need to write the Land function for it to pass.
func TestLand(t *testing.T) {
	ship := NewShip()
	err := ship.Land()
	assert.Nil(t, err)
}
