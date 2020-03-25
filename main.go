package main

import "github.com/yueekee/Neuromancer/initRouter"

func main() {
	router := initRouter.SetupRouter()
	router.Run()
}
