package main

import (
	"github.com/adamluzsi/testcase"
	"testing"
)

func TestHello(t *testing.T) {
	s := testcase.NewSpec(t)

	want := testcase.Var[string]{
		ID: "Hello World!", //what is this suppose to represent ?
		Init: func(t *testcase.T) string {
			return "Hello World!"
		},
	}
	act := func(t *testcase.T) string {
		return HelloWorld()
	}

	s.Describe("HelloWorld", func(s *testcase.Spec) {
		s.When("the user runs HelloWorld(),", func(s *testcase.Spec) {
			s.Then("the function returns Hello World!", func(t *testcase.T) {
				t.Must.Equal(want.Get(t), act(t))
			})
		})
	})

}
