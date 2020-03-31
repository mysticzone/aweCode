package balance

type Balancer interface {
	// DoBalance([]*Instance, ...string) (*Instance, error)
	DoBalance([]*Instance, ...string) (*Instance, error)
}
