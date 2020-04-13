package progress

// Renderer -
type Renderer interface {
	Run()
	Success()
	Fail()
}
