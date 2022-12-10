package main

import "testing"
import "strings"

func TestPrio(t *testing.T) {
	for i, char := range "abcdefghijklmnopqrstuvwxyz" {
		prio := charToPrio(char)
		if prio != i+1 {
			t.Errorf("Invalid %c %d %d ", char, prio, i+1)
		}
	}
	for i, char := range strings.ToUpper("abcdefghijklmnopqrstuvwxyz") {
		prio := charToPrio(char)
		if prio != i+27 {
			t.Errorf("Invalid %c %d %d ", char, prio, i+27)
		}
	}
}
