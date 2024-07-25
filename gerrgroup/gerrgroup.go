package gerrgroup

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type Group struct {
	context.Context
	context.CancelFunc
	gwait   uint32
	wait    *sync.WaitGroup
	errChan chan error
}

func New(bufsize ...int) *Group {
	return NewWithCtx(context.Background(), bufsize...)
}
func NewWithCtx(ctx context.Context, bufsize ...int) *Group {
	chane_size := 10
	if len(bufsize) > 0 {
		chane_size = bufsize[0]
	}
	cancelCtx, cancel := context.WithCancel(ctx)
	return &Group{
		Context:    cancelCtx,
		CancelFunc: cancel,
		wait:       new(sync.WaitGroup),
		errChan:    make(chan error, chane_size),
	}
}

func (g *Group) Go(f func() error) {
	g.GoCtx(func(ctx context.Context) error {
		return f()
	})
}

func (g *Group) GoCtx(f func(context.Context) error) {
	g.wait.Add(1)
	go func() {
		defer func() {
			g.wait.Done()
			if r := recover(); r != nil {
				if re, ok := r.(error); ok {
					g.errChan <- re
				} else {
					g.errChan <- fmt.Errorf("%v", r)
				}
			}
		}()
		select {
		case <-g.Context.Done():
			g.errChan <- g.Context.Err()
		default:
			if err := f(g.Context); err != nil {
				g.errChan <- err
			}
		}
	}()
}

func (g *Group) Wait() <-chan error {
	if atomic.CompareAndSwapUint32(&g.gwait, 0, 1) {
		go func() {
			g.wait.Wait()
			g.CancelFunc()
			close(g.errChan)
		}()
	}
	return g.errChan
}
