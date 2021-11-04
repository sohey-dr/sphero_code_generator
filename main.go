package main

import (
    "sphero_code_generator/router"
)

func main() {
    router := router.InitRouter()
    router.Logger.Fatal(router.Start(":8080"))
}
