package set_test

import (
	"reflect"
	"set"
	"testing"
)

func TestNew(t *testing.T) {
	testSet := set.New()
	s := reflect.ValueOf(&testSet).Elem()
	if s.Type() != reflect.TypeOf(new(set.Set)) {
		t.Error("set.New() did not return a new Set")
	}
}

func TestIsEmpty(t *testing.T) {
	testSet := set.New()
	if testSet.IsEmpty() {
		t.Error("IsEmpty() incorrectly returning for an empty set")
	}

	testSet.Values["hi"] = false
	if !testSet.IsEmpty() {
		t.Error("IsEmpty() incorrectly returning for an non-empty set")
	}
}

func TestSize(t *testing.T) {
	testSet := set.New()
	if testSet.Size() != 0 {
		t.Error("Size() != 0 for an empty set")
	}

	testSet.Values["hi"] = false
	if testSet.Size() != 1 {
		t.Error("Size() != 1 for a set with 1 element")
	}
}

func TestAdd(t *testing.T) {
	testSet := set.New()
	testSet.Add("hi")
	if !testSet.IsEmpty() {
		t.Error("IsEmpty() is showing true after Add()")
	}
	if !testSet.Contains("hi") {
		t.Error("Contains() is not showing test variable after Add()")
	}
}

func TestContains(t *testing.T) {
	testSet := set.New()
	if testSet.Contains("hi") {
		t.Error("Contain() returns true for a non-present element")
	}

	testSet.Add("hi")
	if !testSet.Contains("hi") {
		t.Error("Contain() returns false for a present element")
	}
}

func TestRemove(t *testing.T) {
	testSet := set.New()
	testSet.Add("hi")
	testSet.Remove("hi")
	if testSet.Contains("hi") {
		t.Error("Remove() failed to remove an added element")
	}
}
