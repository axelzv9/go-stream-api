package slice

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestStream_Filter(t *testing.T) {
	s := From([]int{1, 2, 3, 4, 5, 6})
	res := s.Filter(func(i int) bool {
		return i%2 == 0
	}).Collect()

	expected := []int{2, 4, 6}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestStream_Limit(t *testing.T) {
	s := From([]int{1, 2, 3, 4, 5})
	res := s.Limit(3).Collect()

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}

	res2 := s.Limit(10).Collect()
	if !reflect.DeepEqual(res2, []int{1, 2, 3, 4, 5}) {
		t.Errorf("Expected full slice, got %v", res2)
	}
}

func TestStream_Skip(t *testing.T) {
	s := From([]int{1, 2, 3, 4, 5})
	res := s.Skip(2).Collect()

	expected := []int{3, 4, 5}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}

	res2 := s.Skip(10).Collect()
	if len(res2) != 0 {
		t.Errorf("Expected empty slice, got %v", res2)
	}
}

func TestStream_Sorted(t *testing.T) {
	s := From([]int{5, 2, 8, 1, 9})
	res := s.Sorted(func(a, b int) int {
		return a - b
	}).Collect()

	expected := []int{1, 2, 5, 8, 9}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestStream_Distinct(t *testing.T) {
	s := From([]int{1, 2, 2, 3, 4, 4, 5, 1})
	res := s.Distinct().Collect()

	expected := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}

	sEmpty := From([]int{})
	resEmpty := sEmpty.Distinct().Collect()
	if len(resEmpty) != 0 {
		t.Errorf("Expected empty slice, got %v", resEmpty)
	}

	sNoDuplicates := From([]string{"a", "b", "c"})
	resNoDuplicates := sNoDuplicates.Distinct().Collect()
	expectedNoDuplicates := []string{"a", "b", "c"}
	if !reflect.DeepEqual(resNoDuplicates, expectedNoDuplicates) {
		t.Errorf("Expected %v, got %v", expectedNoDuplicates, resNoDuplicates)
	}
}

func TestStream_Reduce(t *testing.T) {
	s := From([]int{1, 2, 3, 4})
	sum := s.Reduce(0, func(a, b int) int {
		return a + b
	})

	if sum != 10 {
		t.Errorf("Expected 10, got %d", sum)
	}
}

func TestStream_Count(t *testing.T) {
	s := From([]int{1, 2, 3})
	if s.Count() != 3 {
		t.Errorf("Expected 3, got %d", s.Count())
	}
}

func TestStream_AnyMatch(t *testing.T) {
	s := From([]int{1, 2, 3})
	if !s.AnyMatch(func(i int) bool { return i == 2 }) {
		t.Error("Expected true")
	}
	if s.AnyMatch(func(i int) bool { return i == 4 }) {
		t.Error("Expected false")
	}
}

func TestStream_AllMatch(t *testing.T) {
	s := From([]int{2, 4, 6})
	if !s.AllMatch(func(i int) bool { return i%2 == 0 }) {
		t.Error("Expected true")
	}
	s2 := From([]int{2, 4, 5})
	if s2.AllMatch(func(i int) bool { return i%2 == 0 }) {
		t.Error("Expected false")
	}
}

func TestStream_FindFirst(t *testing.T) {
	s := From([]int{1, 2, 3})
	val, ok := s.FindFirst()
	if !ok || val != 1 {
		t.Errorf("Expected 1, true; got %v, %v", val, ok)
	}

	sEmpty := From([]int{})
	_, okEmpty := sEmpty.FindFirst()
	if okEmpty {
		t.Error("Expected false for empty stream")
	}
}

func TestMap(t *testing.T) {
	s := From([]int{1, 2, 3})
	res := Map(s, func(i int) string {
		return strconv.Itoa(i)
	}).Collect()

	expected := []string{"1", "2", "3"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestFlatMap(t *testing.T) {
	s := From([]string{"hello", "world"})
	res := FlatMap(s, func(str string) []string {
		return strings.Split(str, "")
	}).Collect()

	expected := []string{"h", "e", "l", "l", "o", "w", "o", "r", "l", "d"}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}

func TestChaining(t *testing.T) {
	s := From([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	res := s.Filter(func(i int) bool { return i%2 == 0 }).
		Skip(1).
		Limit(2).
		Collect()

	// Evens: 2, 4, 6, 8, 10
	// Skip 1: 4, 6, 8, 10
	// Limit 2: 4, 6

	expected := []int{4, 6}
	if !reflect.DeepEqual(res, expected) {
		t.Errorf("Expected %v, got %v", expected, res)
	}
}
