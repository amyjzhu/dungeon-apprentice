package main

import (
	"encoding/json"
	"strings"
	"strconv"
	"fmt"
)

type DamageRoll struct {
	rolls int
	dice int
}

func (dr *DamageRoll) UnmarshalJSON(data []byte) error {
	var roll string

	err := json.Unmarshal(data, &roll);
	if err != nil {
		return err
	}

	parts := strings.Split(roll, "d")
	dr.dice = getInt(parts[0])
	dr.rolls = getInt(parts[1])
	return nil
}

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
