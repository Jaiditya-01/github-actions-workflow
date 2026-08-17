package main

import "testing"

func TestAdd(t *testing.T) {

    result := Add(5, 3)

    expected := 8

    if result != expected {
        t.Errorf("Expected %d but got %d", expected, result)
    }
}