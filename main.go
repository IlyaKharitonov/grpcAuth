package main

import "authService/service"

func main() {
	//добавить обработку паники
	//добавить логирование
	//добавить грейсфул

	service.New().Run()

}
