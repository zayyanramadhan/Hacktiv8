package main

import "webserver/router"

const PORT = ":7000"

func main() {
	router.StartServer().Run(PORT)
}
