package main

import (
	"strconv"
	"fmt"
)



type Action struct {
	name string
	desc string
	attack_bonus int
	damage_dice DamageRoll
	damage_bonus int
}
