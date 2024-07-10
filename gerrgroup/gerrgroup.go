package gerrgroup

import (
	"sync"
)

type Group struct {
	wait    *sync.WaitGroup
	errChan chan error
}

func New() *Group {
	return &Group{
		wait:    new(sync.WaitGroup),
		errChan: make(chan error, 1),
	}
}

func (g *Group) Go(f func() error) {
	g.wait.Add(1)
	go func() {
		defer g.wait.Done()
		if err := f(); err != nil {
			g.errChan <- err
		}
	}()
}

func (g *Group) Wait() <-chan error {
	go func() {
		g.wait.Wait()
		close(g.errChan)
	}()
	return g.errChan
}
