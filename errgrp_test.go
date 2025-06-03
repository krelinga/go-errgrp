package errgrp_test

import (
	"errors"
	"testing"

	"github.com/krelinga/go-errgrp"
)

func TestCut(t *testing.T) {
	g := errgrp.New()
	if g == nil {
		t.Fatal("expected non-nil errgrp")
		return
	}
	if !g.Ok() {
		t.Error("expected errgrp to be ok")
	}
	if g.Join() != nil {
		t.Error("expected errgrp error to be nil")
	}

	err1 := errors.New("error 1")
	err2 := errors.New("error 2")
	func1 := func() (string, error) {
		return "func1_1", err1
	}
	func2 := func() (string, string, error) {
		return "func2_1", "func2_2", err2
	}
	funcOk := func() (string, error) {
		return "ok", nil
	}
	var out1, out2 string

	out1 = errgrp.Cut(funcOk())(g)
	if out1 != "ok" {
		t.Errorf("expected 'ok', got '%s'", out1)
	}
	if g.Join() != nil {
		t.Error("expected errgrp error to be nil after ok function")
	}
	if !g.Ok() {
		t.Error("expected errgrp to be ok after ok function")
	}

	out1 = errgrp.Cut(func1())(g)
	if out1 != "func1_1" {
		t.Errorf("expected 'func1_1', got '%s'", out1)
	}
	if g.Ok() {
		t.Error("expected errgrp to not be ok after error function")
	}
	if !errors.Is(g.Join(), err1) {
		t.Errorf("expected error to be %v, got %v", err1, g.Join())
	}

	out1, out2 = errgrp.Cut2(func2())(g)
	if out1 != "func2_1" {
		t.Errorf("expected 'func2_1', got '%s'", out1)
	}
	if out2 != "func2_2" {
		t.Errorf("expected 'func2_2', got '%s'", out2)
	}
	if g.Ok() {
		t.Error("expected errgrp to not be ok after error function")
	}

	err := g.Join()
	if !errors.Is(err, err1) {
		t.Errorf("expected error to be %v, got %v", err1, err)
	}
	if !errors.Is(err, err2) {
		t.Errorf("expected error to be %v, got %v", err2, err)
	}
}
