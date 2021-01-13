package ufc

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	var dir = "./tmptest"
	if PathExist(dir) != false {
		t.Fatal("fail PathExist")
	}
	if PathNotExist(dir) != true {
		t.Fatal("fail PathNotExist")
	}
	if IsDir(dir) != false {
		t.Fatal("fail IsDir")
	}

	err := CreateDir(dir)
	if err != nil {
		t.Fatalf("create dir fail: %s", dir)
	}
	defer os.Remove(dir)
	if PathExist(dir) != true {
		t.Fatal("after fail PathExist")
	}
	if PathNotExist(dir) != false {
		t.Fatal("after fail PathNotExist")
	}
	if IsDir(dir) != true {
		t.Fatal("after fail IsDir")
	}

	f := "main.go"
	if IsFile(f) != true {
		t.Fatal("fail IsFile")
	}
	if IsFile("/tmp/this_is_a_not_exist_file") != false {
		t.Fatal("fail IsFile No.2")
	}

	shm := "/dev/shm"
	if PathExist(shm) {
		if IsDir(shm) != true {
			t.Fatal("/dev/shm is dir")
		}
	}

	nf := "/dev/null"
	if IsFile(nf) {
		if IsCommonFile(nf) {
			t.Fatal("/dev/null is not common file")
		}
	}

	src := "go.mod"
	dst := "go.mod.bak"
	_, err = FileCopy(dst, src)
	if err != nil {
		t.Log("fail FileCopy")
		t.Fatal(err)
	}
	defer os.Remove(dst)

	if IsFile(dst) != true {
		t.Fatal("created dst, but fail IsFile")
	}
	srcText, err := FileReadByte(src)
	if err != nil {
		t.Fatal("fail FileReadByte")
	}
	dstText, err := FileReadStr(dst)
	if err != nil {
		t.Fatal("fail FileReadStr")
	}
	if string(srcText) != dstText {
		t.Fatal("fail FileCopy, src and dst are inconsistent")
	}

	dstN := "go.mod.bak.N"
	_, err = FileCopyN(dstN, src, 6)
	if err != nil {
		t.Log("fail FileCopyN")
		t.Fatal(err)
	}
	defer os.Remove(dstN)
	dstnText, err := FileReadStr(dstN)
	if err == nil {
		if dstnText != "module" {
			t.Fatal("fail FileCopyN, invalid bytes")
		}
	}
}

func TestBool(t *testing.T) {
	if IsTrue("1") != true {
		t.Fatal("1 is true")
	}
	if IsTrue("t") != true {
		t.Fatal("t is true")
	}
	if IsTrue("T") != true {
		t.Fatal("T is true")
	}
	if IsTrue("true") != true {
		t.Fatal("true is true")
	}
	if IsTrue("True") != true {
		t.Fatal("True is true")
	}
	if IsTrue("TRUE") != true {
		t.Fatal("TRUE is true")
	}
	if IsTrue("abc") == true {
		t.Fatal("abc not true")
	}
	if IsTrue("") == true {
		t.Fatal("empty not true")
	}

	if IsFalse("0") != true {
		t.Fatal("0 is false")
	}
	if IsFalse("f") != true {
		t.Fatal("f is false")
	}
	if IsFalse("F") != true {
		t.Fatal("F is false")
	}
	if IsFalse("false") != true {
		t.Fatal("false is false")
	}
	if IsFalse("False") != true {
		t.Fatal("False is false")
	}
	if IsFalse("FALSE") != true {
		t.Fatal("FALSE is false")
	}
	if IsFalse("abc") != true {
		t.Fatal("abc is false")
	}
	if IsFalse("") != true {
		t.Fatal("empty is false")
	}
}

func TestString(t *testing.T) {
	s := []string{"1", "", "a"}
	if StrInSlice("", s) != true {
		t.Fatal("empty string in slice")
	}
	if StrInSlice("1", s) != true {
		t.Fatal("1 in slice")
	}
	if StrInSlice("a", s) != true {
		t.Fatal("a in slice")
	}
	if StrInSlice("b", s) == true {
		t.Fatal("b not in slice")
	}

	has, index := InArraySlice("", s)
	if has != true {
		t.Fatal("empty string in InArraySlice")
	}
	if index != 1 {
		t.Fatal("empty index error")
	}

	has, _ = InArraySlice("1", s)
	if has != true {
		t.Fatal("1 in InArraySlice")
	}

	has, _ = InArraySlice("a", s)
	if has != true {
		t.Fatal("a in InArraySlice")
	}

	has, _ = InArraySlice("b", s)
	if has == true {
		t.Fatal("b not in InArraySlice")
	}

	has, _ = InArraySlice(1, s)
	if has == true {
		t.Fatal("1(number) not in InArraySlice")
	}

	has, _ = InArraySlice("a", "abc")
	if has == true {
		t.Fatal("a not in abc(string), InArraySlice not allow string")
	}

	has, _ = InArraySlice("a", map[string]int{"a": 1})
	if has == true {
		t.Fatal("a not in map, InArraySlice not allow map")
	}

	var s2 [][]string
	s2 = append(s2, s)
	s2 = append(s2, []string{"b"})
	s2 = append(s2, []string{"1"})

	has, _ = InArraySlice(s, s2)
	if has != true {
		t.Fatal("s in s2")
	}

	has, _ = InArraySlice([]string{"b"}, s2)
	if has != true {
		t.Fatal("[]string{b} in s2")
	}

	has, _ = InArraySlice([]string{"c"}, s2)
	if has == true {
		t.Fatal("[]string{c} not in s2")
	}

	has, _ = InArraySlice([]string{"1"}, s2)
	if has != true {
		t.Fatal("[]string{1} in s2")
	}

	has, _ = InArraySlice([]int{1}, s2)
	if has == true {
		t.Fatal("[]int{1} not in s2")
	}
}
