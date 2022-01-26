package spaceship

import (
	"errors"
	"log"
)

type Bullet struct {
	Size     int
	Distance int
}

type SimpleChassis struct {
	Engine      Engine
	Lights      interface{} // leaving as empty interface for itme being while we design those systems.
	Ammunition  []Bullet
	LandingGear []interface{}
	// TODO: Add a weapons array e.g. []Weapon
	// TODO: Add a fuel gauge
	// TODO: Add a distance to earth metric
}

type Ship interface {
	Takeoff() error
	Land() error
}

// Force SimpleChassis to fulfill Ship
type Engine interface {
	Run() error
}

// Question: Should we move engine stuff out into a different file?

// FusionEngine fulfills Engine
type FusionEngine struct{}

var FusionEngine _ = (Engine)(nil)

// UpgradedFusion fulfills engine at runtime
var UpgradedFusion _ = (Engine)(nil)

// Force SimpleChassis to fulfill Ship
var SimpleChassis _ = (Ship)(nil)

// spec 1: write your constructor function
func NewShip() Ship {
	sc := SimpleChassis{
		// TODO: update to upgraded fusion engine library
		Engine: UpgradedFusionEngine{},
	}
}

func (c *SimpleChassis) SetEngine(e Engine) error {
	c.Engine = e
}

func (c *Chassis) Takeoff() error {
	err := c.Engine.Run()
	if err != nil {
		return err
	}
	log.Printf("TAKING OFF")
	return nil
}

func (c *Chassis) Land() error {
	return errors.New("NOT IMPLEMENTED")
}
