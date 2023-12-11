package main

import (
	"kawa/gradingservice/internal/app/grading"
)

func main() {
	app := grading.NewApp()
	app.Run()
}
