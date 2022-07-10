package prott

import (
	"fmt"
)

type Dispatch[T comparable, D any] struct {
	handlers map[T][]func(data D)
}

func (d *Dispatch[T, D]) Init() {
	d.handlers = make(map[T][]func(data D))
}

func (d *Dispatch[T, D]) Register(t T, f func(data D)) {
	d.ensureArrayExists(t)

	d.handlers[t] = append(d.handlers[t], f)
}

func (d *Dispatch[T, D]) ensureArrayExists(t T) {
	_, ok := d.handlers[t]
	if !ok {
		d.handlers[t] = make([]func(data D), 0)
	}
}

func (d *Dispatch[T, D]) Disp(t T, data D) error {
	functions, ok := d.handlers[t]
	if !ok {
		return fmt.Errorf("type has not been registered")
	}

	for _, fn := range functions {
		fn(data)
	}

	return nil
}
