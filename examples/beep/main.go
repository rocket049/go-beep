package main

import (
	"gitee.com/rocket049/go-beep"
)

func main() {
	player, err := beep.NewBeepPlayer()
	if err != nil {
		panic(err)
	}
	defer player.Close()
	snds := []int{500, 600, 700, 800, 900, 1000, 1100, 1200, 1300, 1400, 1500, 1600, 1700, 1800, 1900, 2000}
	for _, v := range snds {
		player.Beep(v, 500)
	}
}
