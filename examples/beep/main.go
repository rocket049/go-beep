package main

import (
	"time"

	"gitee.com/rocket049/go-beep"
)

func main() {
	player, err := beep.NewBeepPlayer()
	if err != nil {
		panic(err)
	}
	defer player.Close()
	snds := []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000}
	player.Beep(0, 1000)
	for _, v := range snds {
		player.Beep(v, 500)
	}

	player.Beep(0, 1000)
	time.Sleep(time.Second * 3)

}
