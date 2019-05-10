package main

import (
	"testing"

	"github.com/powerman/check"
)

func TestHello(tt *testing.T) {
	t := check.T(tt)

	t.Equal(hello(), "hello world")
}
