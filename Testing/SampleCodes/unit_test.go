package main

import testing

func Add(a, b int) int {
    return a + b
}

// add_test.go
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("expected 5, got %d", result)
    }
}