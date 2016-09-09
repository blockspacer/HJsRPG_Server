package main

import (
    "fmt"
    gs "GameServer"
)

func main(){
    gameService := gs.CreateGameService()
    fmt.Println(gameService.ID)

    gameService.Start()
    gameService.WriteLog(fmt.Sprintf("Test Log(%s, %d, %f)", "T1", 123, 1.23))
}