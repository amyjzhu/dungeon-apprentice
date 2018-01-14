package main

import (
	"strconv"
	"fmt"
)


func getInt(s string) int {
	integer, err := strconv.Atoi(s)
	if err == nil {
		fmt.Println("Talk to Amy... %s", err)
	}

	return integer
}


type Action struct {
	name string
	desc string
	attack_bonus int
	damage_dice DamageRoll
	damage_bonus int
}
