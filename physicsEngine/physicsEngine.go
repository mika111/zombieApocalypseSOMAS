package physicsEngine

type Vector2D struct {
	X float32
	Y float32
}

func (v2d *Vector2D) Add(vec Vector2D) {
	v2d.X += vec.X
	v2d.Y += vec.Y
}

func (v2d *Vector2D) Div(k float32) Vector2D {
	return Vector2D{
		X: v2d.X / k,
		Y: v2d.Y / k,
	}
}

type PhysicalState struct {
	Position Vector2D
	Velocity Vector2D
	Mass     float32
}

func (ps *PhysicalState) UpdatePhysicalState(force Vector2D) {
	acc := force.Div(ps.Mass)
	ps.Velocity.Add(acc)
	ps.Position.Add(ps.Velocity)
}

type IPhysicsObject interface {
	UpdatePhysicalState(Vector2D)
	GetPhysicalState() PhysicalState
}
