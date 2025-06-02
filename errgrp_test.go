package errgrp_test

import (
	"errors"
	"testing"

	errgrp "github.com/krelinga/go-errorgroup"
	"github.com/stretchr/testify/assert"
)

func TestCut(t *testing.T) {
	g := errgrp.New()
	if !assert.NotNil(t, g) {
		return
	}
	assert.True(t, g.Ok())
	assert.NoError(t, g.Join())

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
	out1 = errgrp.Cut(func1())(g)
	assert.Equal(t, "func1_1", out1)
	out1, out2 = errgrp.Cut2(func2())(g)
	assert.Equal(t, "func2_1", out1)
	assert.Equal(t, "func2_2", out2)
	out1 = errgrp.Cut(funcOk())(g)
	assert.Equal(t, "ok", out1)
	err := g.Join()
	assert.True(t, errors.Is(err, err1))
	assert.True(t, errors.Is(err, err2))
}
