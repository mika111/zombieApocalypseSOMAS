package physicsengine

type Vector2D struct {
	X float32
	Y float32
}

type PhysicalState struct {
	Position Vector2D
	Velocity Vector2D
	Mass float32
}

type IPhysicsObject interface {
	UpdatePhysicalState(Vector2D)
	GetPhysicalState() PhysicalState
}

