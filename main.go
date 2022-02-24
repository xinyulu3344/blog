package main

import (
    "blog/model"
    "blog/routes"
)

func main() {
    model.InitDB()
    routes.InitRouter()
}
