package errgrp

import "errors"

type Group struct {
	errs []error
	ok   bool
}

func New() *Group {
	return &Group{
		errs: make([]error, 0),
		ok:   true,
	}
}

func (g *Group) Add(err error) {
	if err != nil {
		g.errs = append(g.errs, err)
		g.ok = false
	}
}

func (g *Group) Join() error {
	if g.ok {
		return nil
	}
	return errors.Join(g.errs...)
}

func (g *Group) Ok() bool {
	return g.ok
}

func Cut[T any](v T, err error) func(*Group) T {
	return func(g *Group) T {
		g.Add(err)
		return v
	}
}

func Cut2[T1, T2 any](v1 T1, v2 T2, err error) func(*Group) (T1, T2) {
	return func(g *Group) (T1, T2) {
		g.Add(err)
		return v1, v2
	}
}
