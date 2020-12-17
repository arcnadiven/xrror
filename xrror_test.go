package xrror

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	SetPathLayer(3)
}

func TestNewXrror(t *testing.T) {
	if err := CatchError(); err != nil {
		fmt.Println(err)
	}
}

func CatchError() error {
	_, err := os.Open("dsdasd")
	if err != nil {
		return New(err)
	}
	return nil
}
