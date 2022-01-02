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
	if err := catchError(); err != nil {
		fmt.Println(err)
	}
}

func catchError() error {
	_, err := os.Open("dsdasd")
	if err != nil {
		return WithStack(err)
	}
	return nil
}
