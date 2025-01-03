package controllers

import (
	"sync"
)

// Wrapper equivalent in Go
type Wrapper[T any] struct {
	Value T
}

// SettingTask equivalent in Go
type SettingTask struct {
	w *Wrapper[string]
	s string
}

func NewSettingTask(w *Wrapper[string], s string) *SettingTask {
	return &SettingTask{w: w, s: s}
}

func (st *SettingTask) Run() {
	st.w.Value = st.s
}

// SwitchingTask equivalent in Go
type SwitchingTask struct {
	w  *Wrapper[string]
	s  string
	mu sync.Mutex
}

func NewSwitchingTask(w *Wrapper[string], s string) *SwitchingTask {
	return &SwitchingTask{w: w, s: s}
}

func (st *SwitchingTask) Run() {
	st.mu.Lock()
	defer st.mu.Unlock()
	if st.w.Value == "" {
		st.w.Value = st.s
	} else {
		st.w.Value = ""
	}
}

// NullAndRestore equivalent in Go
type NullAndRestore struct {
	s        string
	original string
	mu       sync.Mutex
	cond     *sync.Cond
}

func NewNullAndRestore(s string) *NullAndRestore {
	nr := &NullAndRestore{
		s:        s,
		original: s,
	}
	nr.cond = sync.NewCond(&nr.mu)
	return nr
}

func (nr *NullAndRestore) NullMethod() {
	nr.mu.Lock()
	defer nr.mu.Unlock()
	nr.s = ""
	nr.cond.Broadcast()
}

func (nr *NullAndRestore) Restore() {
	nr.mu.Lock()
	defer nr.mu.Unlock()
	for nr.s != "" {
		nr.cond.Wait()
	}
	nr.s = nr.original
}

func (nr *NullAndRestore) Get() string {
	nr.mu.Lock()
	defer nr.mu.Unlock()
	return nr.s
}
