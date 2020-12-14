package plugins

import (
	"fmt"
	"os"
	"plugin"
	"testing"
)

// TestPluginRandInt ensures that the go plugin is actually generating random integers
func TestPluginRandInt(t *testing.T) {
	plug, err := plugin.Open("./plugin.so")
	if err != nil {
		panic(err)
	}

	randInt, err := plug.Lookup("RandInt")
	if err != nil {
		panic(err)
	}

	randFunc, ok := randInt.(func() int)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	randNum := randFunc()
	if randNum < 0 {
		t.Errorf("expected a random integer but received %d", randNum)
	}

	randNum2 := randFunc()
	if randNum == randNum2 {
		t.Errorf("expected 2 random integers but both integers are the same %d = %d", randNum, randNum2)
	}
}

// BenchmarkRandInt tests math/rand for generating random integers without a go plugin
func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt()
	}
}

// BenchmarkPluginRandInt uses a go plugin and tests math/rand for generating random integers
func BenchmarkPluginRandInt(b *testing.B) {
	plug, err := plugin.Open("./plugin.so")
	if err != nil {
		panic(err)
	}

	randInt, err := plug.Lookup("RandInt")
	if err != nil {
		panic(err)
	}

	randFunc, ok := randInt.(func() int)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		randFunc()
	}
}
