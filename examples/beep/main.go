package main

import (
	"fmt"
	"time"

	"gitee.com/rocket049/go-beep"
)

func main() {
	player, err := beep.NewBeepPlayer()
	if err != nil {
		panic(err)
	}
	defer player.Close()
	snds := []int{350, 400, 500, 600, 700, 800, 900}
	player.Beep(0, 1000)
	for _, v := range snds {
		fmt.Printf("Freq:%d\n", v)
		player.Beep(v, 1000)

	}

	time.Sleep(time.Second * 2)

}
