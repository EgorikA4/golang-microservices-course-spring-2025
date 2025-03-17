package main

import "golang-microservices-course-spring-2025/internal/application"

func main() {
	app := application.NewApp()
	app.Listen(":6080")
}
