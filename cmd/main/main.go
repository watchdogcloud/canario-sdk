package main

import "github.com/zakhaev26/canario/pkg/canario"

func main() {
	c := canario.NewCanario()
	c.RunPeriodicMetrics()

}
