package physicsengine

type Force struct {
	XComponent float32
	YComponent float32 
}

type Position struct {
	XComponent float32
	YComponent float32
}

type PhysicalState struct {
	Force Force
	Position Position
}

type IPhysicsObject interface {
	UpdatePhysicalState(PhysicalState) PhysicalState
}