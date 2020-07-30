package aefire

import "sync"

type KeyMutax struct {
	sync.Map
}

func (v *KeyMutax) Lock(key string) {
	m, _ := v.LoadOrStore(key, &sync.Mutex{})
	m.(*sync.Mutex).Lock()
}

func (v *KeyMutax) UnLock(key string) {
	m, ok := v.Load(key)
	if ok {
		m.(*sync.Mutex).Unlock()
	}
}
