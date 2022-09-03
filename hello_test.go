package main

import (
	"github.com/adamluzsi/testcase"
	"testing"
)

func TestHello(t *testing.T) {
	s := testcase.NewSpec(t)
	var (
		want = "Hello World!"
		act  = func(t *testcase.T) string {
			return HelloWorld()
		}
	)

	s.Test("When I run HelloWorld", func(t *testcase.T) {
		t.Must.Equal(want, act(t))
	})

}
