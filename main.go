package main

import "github.com/braddle/hex-arch-example-go/app"

func main() {
	m := app.NewMicroservice()
	m.Serve(":8080")
}
