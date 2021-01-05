package ufc

import (
	"os"
	"testing"
)

func TestFuncs(t *testing.T) {
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
		t.FailNow()
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

}
