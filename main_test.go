package plugins

import (
	"fmt"
	"math/rand"
	"os"
	"plugin"
	"testing"
)

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}

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
