package balance

import "errors"

func init() {
	RegisterBalance("robin", &RobinBalance{})
}

type RobinBalance struct {
	curIndex int
}

func (p *RobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {

	if len(insts) == 0 {
		err = errors.New("No instances")
		return
	}
	lens := len(insts)
	if p.curIndex >= lens {
		p.curIndex = 0
	}
	inst = insts[p.curIndex]
	p.curIndex = (p.curIndex + 1) % lens
	return
}
