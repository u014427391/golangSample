package main

import (
	"testing"
)


func TestAdd(t *testing.T)  {
	res := add(1,2)
	if res != 3 {
		t.Fatalf("error res %v" , res)
	}
	t.Logf("success res %v" , res)
}

func TestAddUpper(t *testing.T)  {
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("error res %v" , res)
	}
	t.Logf("success res %v" , res)
}

