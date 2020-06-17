// package game implements simple primitives for building interactive 2D games
package game

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
)

type Vec2 struct {
	float64 X
	float64 Y
}

type worldRenderer struct {
	world *GameWorld
}

// GameWorld represents the top-level game object, and also implement a Fyne
// widget which renders the game.
type GameWorld struct {
}

type Entity interface {
	Objects() []fyne.CanvasObject
	Position() Vec2
	PhysicsTick(deltat float64)
	PhysicsImpulse(Fx, Fy float64)
}

type Player struct {
	pos     Vec2
	vel     Vec2
	heading Vec2
	sprite  canvas.Image
}

func NewPlayer(spritePath string) (*Player, error) {
	p := &Player{}

	p.sprite = NewImageFromFile(spritePath)
}

func (p *Player) Objects() []fyne.CanvasObject {
	obj := make([]fyne.CanvasObject, 1)

	return obj
}
