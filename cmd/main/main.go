package main

import "github.com/zakhaev26/canario/pkg/canario"

func main() {

	cio := canario.NewCanario()
	cio.RunPeriodicMetrics()
}
