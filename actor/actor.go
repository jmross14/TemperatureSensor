// Package actor is used to define a common interface for Actors
package actor

// Actor is a common interface for actor behavior
type Actor interface {
	StartActor() *Actor
	Tell()
	Ask(chan struct{})
}

