package balance

import "fmt"

type BalanceMgr struct {
	allBalancer map[string]Balancer
}

var (
	mgr = BalanceMgr{
		allBalancer: make(map[string]Balancer),
	}
)

func (p *BalanceMgr) registerBalance(name string, b Balancer) {
	p.allBalancer[name] = b
}

func RegisterBalance(name string, b Balancer) {
	mgr.registerBalance(name, b)
}

func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	balance, ok := mgr.allBalancer[name]
	if !ok {
		fmt.Errorf("Not found %s balancer", name)
		return
	}
	fmt.Printf("use %s balance\n", name)
	inst, err = balance.DoBalance(insts)
	return
}
