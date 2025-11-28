package ptr

import (
	"testing"
)

func TestRef(t *testing.T) {
	val := 42
	ptr := Ref(val)

	if ptr == nil {
		t.Error("Expected non-nil pointer")
	}
	if *ptr != val {
		t.Errorf("Expected %d, got %d", val, *ptr)
	}

	str := "hello"
	strPtr := Ref(str)
	if *strPtr != str {
		t.Errorf("Expected %s, got %s", str, *strPtr)
	}
}

func TestDefaultIfNil(t *testing.T) {
	var nilPtr *int
	result := DefaultIfNil(nilPtr, 100)
	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}

	val := 42
	ptr := &val
	result2 := DefaultIfNil(ptr, 100)
	if result2 != 42 {
		t.Errorf("Expected 42, got %d", result2)
	}

	var nilStrPtr *string
	strResult := DefaultIfNil(nilStrPtr, "default")
	if strResult != "default" {
		t.Errorf("Expected 'default', got %s", strResult)
	}
}

func TestDefaultIfZero(t *testing.T) {
	result := DefaultIfZero(0, 100)
	if result != 100 {
		t.Errorf("Expected 100, got %d", result)
	}

	result2 := DefaultIfZero(42, 100)
	if result2 != 42 {
		t.Errorf("Expected 42, got %d", result2)
	}

	strResult := DefaultIfZero("", "default")
	if strResult != "default" {
		t.Errorf("Expected 'default', got %s", strResult)
	}

	strResult2 := DefaultIfZero("hello", "default")
	if strResult2 != "hello" {
		t.Errorf("Expected 'hello', got %s", strResult2)
	}

	boolResult := DefaultIfZero(false, true)
	if boolResult != true {
		t.Errorf("Expected true, got %v", boolResult)
	}

	boolResult2 := DefaultIfZero(true, false)
	if boolResult2 != true {
		t.Errorf("Expected true, got %v", boolResult2)
	}
}
