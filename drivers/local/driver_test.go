package local

import (
	"fmt"
	"testing"
)

var localDriver = new(Local)

func TestLocalDriverInit(t *testing.T) {
	err := localDriver.Init(`{"path":"/tmp","show_hidden":true}`)
	if err != nil {
		t.Fatal(err)
	}

	if localDriver.Path != `/tmp` {
		t.Fatalf(`Expected path to be /tmp, got %s`, localDriver.Path)
	}

	if localDriver.ShowHidden != true {
		t.Fatalf(`Expected show_hidden to be true, got %t`, localDriver.ShowHidden)
	}
}

func TestLocalDriverList(t *testing.T) {
	err := localDriver.Init(`{"path":"/tmp","show_hidden":false}`)
	if err != nil {
		t.Fatal(err)
	}

	list, err := localDriver.List("/Users/zxdstyle")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(list)
}
