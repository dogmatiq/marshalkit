package portablename_test

import (
	"reflect"
	"testing"

	. "github.com/dogmatiq/enginekit/enginetest/stubs"
	. "github.com/dogmatiq/marshalkit/codec/internal/portablename"
)

type MultipleTypeParameters[A, B any] struct{}

func TestFromReflect(t *testing.T) {
	cases := []struct {
		Type reflect.Type
		Want string
	}{
		{reflect.TypeFor[*AggregateRootStub](), "AggregateRootStub"},
		{reflect.TypeFor[CommandStub[TypeA]](), "CommandStub[TypeA]"},
		{reflect.TypeFor[CommandStub[int]](), "CommandStub[int]"},
		{reflect.TypeFor[MultipleTypeParameters[TypeA, TypeB]](), "MultipleTypeParameters[TypeA,TypeB]"},
	}

	for _, c := range cases {
		got, ok := FromReflect(c.Type)
		if !ok {
			t.Fatal("expected portable name to be valid")
		}

		if got != c.Want {
			t.Errorf("unexpected portable name for %s: got %q, want %q", c.Type, got, c.Want)
		}
	}

	_, ok := FromReflect(reflect.TypeOf(struct{}{}))
	if ok {
		t.Fatal("expected portable name to be invalid for anonymous type")
	}
}
