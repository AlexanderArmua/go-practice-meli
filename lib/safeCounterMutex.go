package lib

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	V   map[string]int
	mux sync.Mutex
}

func (c *SafeCounter) Incrementar(key string, safe bool) {
	if safe {
		c.mux.Lock()
		defer c.mux.Unlock()
	}

	// Incrementamos y nos aseguramos que solo una goroutine tenga acceso al recurso
	c.V[key]++
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.V[key]
}

func TestMutex() {
	c := SafeCounter{V: make(map[string]int)}
	iteraciones := 10000
	safe := true

	for i := 0; i < iteraciones; i++ {
		go c.Incrementar("algo", safe)
	}

	time.Sleep(time.Second * 5)

	fmt.Printf("Iteraciones: %v - Safe: %v - Resultado: %v.\n", iteraciones, safe, c.Value("algo"))
}
