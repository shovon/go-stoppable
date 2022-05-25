package stoppable

import (
	"sync"
	"testing"
	"time"
)

type TestStoppable struct {
	Stoppable
}

func TestStop(t *testing.T) {
	v := TestStoppable{NewStoppable()}
	test(&v, t)
}

func test(v *TestStoppable, t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for !v.HasStopped() {
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		<-v.OnStopped()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if v.HasStopped() {
			t.Fail()
		}
		for i := 0; i < 100; i++ {
			<-time.After(time.Nanosecond * 100)
		}
		v.Stop()
		wg.Done()
	}()

	wg.Wait()
}

func TestRestart(t *testing.T) {
	v := TestStoppable{NewStoppable()}

	test(&v, t)
	v.Stoppable = NewStoppable()

	test(&v, t)
	v.Stoppable = NewStoppable()

	test(&v, t)
	v.Stoppable = NewStoppable()
}
