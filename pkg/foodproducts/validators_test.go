package foodproducts

import (
	"testing"
	"time"
)

func TestFoo(t *testing.T) {
	time.Sleep(2 * time.Second)
	t.Fail()
}

func TestBar(t *testing.T) {
	t.Skip()
}
