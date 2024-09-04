package task

import (
	"github.com/i-nas/i-nas-tools/pkg/lifecycle"
	"github.com/panjf2000/ants/v2"
)

var p *ants.Pool

func init() {
	var err error
	p, err = ants.NewPool(100)
	if err != nil {
		panic(err)
	}
	lifecycle.RegisterBeforeExit(
		func() {
			p.Release()
		},
	)
}

func Submit(f func()) {
	p.Submit(f)
}

func SubmitAwait(f func()) {
	ch := make(chan struct{})
	p.Submit(func() {
		f()
		close(ch)
	})
	<-ch
}
