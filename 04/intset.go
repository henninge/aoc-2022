package main

import "fmt"

type Set map[int]bool

func (s *Set) Add(val int) bool {
	if _, ok := (*s)[val]; ok {
		return false
	}
	(*s)[val] = true
	return true
}

func (s *Set) AddRange(val1, val2 int) {
	for i := val1; i <= val2; i++ {
		s.Add(i)
	}
}

func (s *Set) IsIn(val int) bool {
	_, ok := (*s)[val]
	return ok
}

func (s *Set) Contains(other *Set) bool {
	for val := range *other {
		if !s.IsIn(val) {
			return false
		}
	}
	return true
}

func (s *Set) Intersects(other *Set) bool {
	for val := range *s {
		if other.IsIn(val) {
			return true
		}
	}
	return false
}

func (s *Set) AsString() string {
	for val := range *s {
		return fmt.Sprintf("[%d]", val)
	}
	return "[]"
}
