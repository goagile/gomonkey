package token

import (
	"testing"
)

func Test_Eq_true(t *testing.T) {

	want := true

	a := New("=", ASSIGN)
	b := New("=", ASSIGN)

	got := a.Equal(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}

}

func Test_Eq_false(t *testing.T) {

	want := false

	a := New("=", ASSIGN)
	b := New("+", PLUS)

	got := a.Equal(b)

	if want != got {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}

}