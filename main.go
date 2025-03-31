package main

import "authService/service"

func main() {
	service := service.New()

	go service.Run()

	service.Stop()
}
