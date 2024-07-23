package main

import (
	"errors"
	"fmt"
)

type Set struct {
	Elements map[string]struct{}
}

func NewSet() *Set {
	var set Set
	set.Elements = make(map[string]struct{})
	return &set
}

func (s *Set) Add(elem string) {
	s.Elements[elem] = struct{}{}
}

func (s *Set) Delete(elem string) error {
	if _, exists := s.Elements[elem]; !exists {
		return errors.New("element doesn't exist in set")
	}
	delete(s.Elements, elem)
	return nil
}

func (s *Set) Contains(elem string) bool {
	_, exists := s.Elements[elem]
	return exists
}

func (s *Set) List() {
	for key := range s.Elements {
		fmt.Println(key)
	}
}

func main() {
	set := NewSet()
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, v := range arr {
		set.Add(v)
	}
	set.List()
}
