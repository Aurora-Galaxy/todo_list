package main

import (
    "todo_list/conf"
    "todo_list/routes"
)

// @title ToDoList API
// @version 0.0.1
// @description This is a sample Server pets
// @name LLL
// @BasePath /api/v1
func main() {
    conf.Init()
    r := routes.NewRouter()
    r.Run(conf.HttpPort)
}
