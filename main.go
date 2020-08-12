package main

import (
	"log"
	"os"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)
	app := App{}
	app.Initialize()
	app.Run()
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
