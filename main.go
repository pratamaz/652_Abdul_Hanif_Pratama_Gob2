package main

import "belajar-gin/routers"

func main() {
	var PORT = ":8082"

	routers.StartServer().Run(PORT)
}