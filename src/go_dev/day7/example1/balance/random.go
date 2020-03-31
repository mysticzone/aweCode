package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalance("random", &RandBalance{})
}

type RandBalance struct {
}

func (p *RandBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {

	if len(insts) == 0 {
		err = errors.New("No instances")
		return
	}

	lens := len(insts)
	index := rand.Intn(lens)
	inst = insts[index]
	return
}
