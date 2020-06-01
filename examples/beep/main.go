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
	snds := []int{700, 1000, 1500, 2000}
	for _, v := range snds {
		player.Beep(v, 300)
	}
}
