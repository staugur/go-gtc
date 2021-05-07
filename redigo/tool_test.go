package redigo

import "testing"

func TestTool(t *testing.T) {
	s := []string{"1", "", "a"}
	if inSlice("", s) != true {
		t.Fatal("empty string in slice")
	}
	if inSlice("1", s) != true {
		t.Fatal("1 in slice")
	}
	if inSlice("a", s) != true {
		t.Fatal("a in slice")
	}
	if inSlice("b", s) == true {
		t.Fatal("b not in slice")
	}

	if KPV("a", []string{"b", "c"})[0].(string) != "a" {
		t.Fatal("kpv error")
	}
}
