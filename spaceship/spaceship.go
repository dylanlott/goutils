package spaceship

import (
	"errors"
	"log"
)

type Bullet struct {
	Size     int
	Distance int
}

// TODO: Write a comment for SimpleChassis and explain what all of these systems we've attached
// to the chassis like Engine and LandingGear do and, sometimes more importantly, what they can't do.
type SimpleChassis struct {
	Engine      Engine
	Lights      interface{} // empty interface{} types are nice for mocking out a design because they can be anything
	Ammunition  []Bullet
	LandingGear []interface{}
	// TODO: Add a weapons array of either a map or a slice of weapons and tell me why you chose map or slice
	// TODO: Add a fuel gauge. You'll need to add a property on the SimpleChassis
	// struct or the Engine interface and then I would write a function to return that value when called.
	// QUESTION: Should the fuel gauge be on the Ship or the Engine interface? Why?
	// ANSWER: TODO
	// TODO: Add an inventory system. This ship will inevitably need to carry inventory, so we need a way to track that.
	// QUESTION: What does the inventory system need to do? Describe your inventory system here and how we would
	// model that in code as part of the SimpleChassis.
	// ANSWER: TODO
	// SPEC 5
	// TODO: Add a messages array for the captain to push messages into.
	// QUESTION: What does this need to look like? What do messages need to track? How should you add a message to the system?
	// ANSWER: TODO
}

// QUESTION: Should we move our interfaces into a different file?
// ANSWER: TODO
type Ship interface {
	Takeoff() error
	Land() error

	// TODO: Write a function that your SimpleChassis struct fulfills that will return the ship's ID.
	// NOTE: These types of functions are often called "getters": They don't take an argument
	// and return a value.
	ID() string // TODO: You'll need to make your SimpleChassis function have this function or else your code won't even compile.
}

// Force SimpleChassis to fulfill Ship
type Engine interface {
	// QUESTION: What if we wanted to change this function from Run to Start.
	// TODO: Change this function from Run to Start
	Run() error
	// TODO: Add a SetThrottle function here that lets us set the Engine's throttle.
	// QUESTION: How should we track the throttle? How would you model that in code?
	// ANSWER: TODO
}

// QUESTION: Should we move engine stuff out into a different file?
// ANSWER: TODO

// FusionEngine fulfills Engine
type FusionEngine struct{}

// UpgradedFusionEngine is our bigger, badder, hella more doper engine.
type UpgradedFusionEngine struct {
	FuelType string
}

func (ufg *UpgradedFusionEngine) Run() error {
	log.Printf("OH YEAH, WE'RE CRUISING NOW, HOSS")
	return nil
}

//
func (e *FusionEngine) Run() error {
	log.Printf("fusion engine running")
	return nil
}

// Force SimpleChassis to fulfill Ship
var _ Ship = (*SimpleChassis)(nil) // NOTE: This line will produce a compile time error until you write that ID function.

// TODO: Write the comment for NewShip and explain what it does.
// NOTE: Functions like this are used for setting up your library and starting any processes that your code depends on.
func NewShip() *SimpleChassis {
	sc := &SimpleChassis{
		Engine: &FusionEngine{}, // QUESTION: I made this a pointer - why do you think it should be a pointer? ANSWER: TODO
	}
	return sc
}

// SetEngine is a function for swapping the engine of a ship. This is meant to be an example of interface fulfillment
// at runtime - a fancy way of saying you could drastically change the behavior of a system or component like this
// with a simple function call.
func (c *SimpleChassis) SetEngine(e Engine) error {
	c.Engine = e
	// NOTE: Sometimes, with interfaces like this, you might never actually have an error to return,
	// but the interface might allow you to return one, so you'll have code like this where you just
	// always return nil.
	return nil
}

// TODO: Write the comment for this function
func (c *SimpleChassis) Takeoff() error {
	err := c.Engine.Run()
	if err != nil {
		return err
	}
	log.Printf("TAKING OFF")
	return nil
}

// Land prepares our ship for landing and then starts to land.
func (c *SimpleChassis) Land() error {
	// SPEC 3
	// TODO: write the Land function.
	// QUESTIONS: What broad steps should we take to land?
	// Do we need to log anything? Do we need to change any properties on our SimpleChassis struct?
	// ANSWER: TODO
	return errors.New("NOT IMPLEMENTED")
}

func (c *SimpleChassis) ID() string {
	// SPEC 2
	// TODO: You'll need to write this function
	return ""
}

// TODO: Write an inventory interface that your inventory system fulfills.
// This should be one of the later things you do.
