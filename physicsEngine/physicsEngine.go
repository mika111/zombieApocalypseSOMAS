package physicsEngine

import (
	"fmt"
)

type Vector2D struct {
	X int
	Y int
}

func ZeroVector() Vector2D {
	return Vector2D{X: 0, Y: 0}
}

func (v2d *Vector2D) Add(vec Vector2D) {
	v2d.X += vec.X
	v2d.Y += vec.Y
}

func (v2d *Vector2D) Div(k int) Vector2D {
	return Vector2D{
		X: v2d.X / k,
		Y: v2d.Y / k,
	}
}

func (v2d *Vector2D) Print() {
	fmt.Printf("[%v,%v]", v2d.X, v2d.Y)
}

func MakeVec2D(X, Y int) Vector2D {
	return Vector2D{X: X,
		Y: Y}
}

type PhysicalState struct {
	Position *Vector2D
	Velocity *Vector2D
	Mass     int
}

func (ps *PhysicalState) UpdatePhysicalState(force Vector2D) {
	acc := force.Div(ps.Mass)
	ps.Velocity.Add(acc)
	ps.Position.Add(*ps.Velocity)
}

func (ps *PhysicalState) GetPhysicalState() PhysicalState {
	return PhysicalState{
		Position: ps.Position,
		Velocity: ps.Velocity,
		Mass:     ps.Mass,
	}
}

func (ps *PhysicalState) PrintPhysicalState() {
	fmt.Printf("Position: %v. Velocity: %v. Mass: %v\n", ps.Position, ps.Velocity, ps.Mass)
}

type IPhysicsObject interface {
	UpdatePhysicalState(Vector2D)
	GetPhysicalState() PhysicalState
	PrintPhysicalState()
}

func CreateInitialPhysicalState(initialPosition *Vector2D, mass int) *PhysicalState {
	return &PhysicalState{Position: initialPosition, Velocity: &Vector2D{X: 0, Y: 0}, Mass: mass}
}
