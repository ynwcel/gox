package gerrgroup

import (
	"errors"
	"testing"
	"time"
)

func TestGroup0(t *testing.T) {
	g := New()
	t.Log(<-g.Wait())
}
func TestGroup(t *testing.T) {
	g := New()
	g.Go(func() error {
		time.Sleep(1 * time.Second)
		t.Log("1s finish")
		return nil
	})
	g.Go(func() error {
		time.Sleep(2 * time.Second)
		t.Log("2s finish")
		return nil
	})
	t.Log(<-g.Wait())
}

func TestGroup1(t *testing.T) {
	g := New()
	g.Go(func() error {
		time.Sleep(3 * time.Second)
		t.Log("group1 3s finish")
		return nil
	})
	g.Go(func() error {
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()
		idx := 0
		for range ticker.C {
			idx += 1
			if idx >= 10 {
				return errors.New("group1.ticker.idx>=3")
			}
		}
		return nil
	})
	t.Log("wait-error:", <-g.Wait())
}
