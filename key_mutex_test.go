package aefire

import (
	"sync"
	"testing"
)

func TestKeyMutax_Lock(t *testing.T) {
	km := KeyMutax{}

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		defer wg.Done()
		km.Lock("a")
		println("a")
	}()
	go func() {
		defer wg.Done()
		km.Lock("b")
		println("b")
		km.UnLock("b")
	}()
	go func() {
		defer wg.Done()
		km.Lock("a")
		println("a1")
	}()
	go func() {
		defer wg.Done()
		km.Lock("b")
		println("b1")
		km.UnLock("b")
	}()

	wg.Wait()
}
