package main

import "fmt"

const MAX_FOOD = 100

type dog struct {
	name string
	food int
}

func (d dog) isHappy() bool {
	if d.food > MAX_FOOD/2 {
		return true
	}
	return false
}

type iHappy interface {
	isHappy() bool
}

func getHappy(many []iHappy) []iHappy {
	h := make([]iHappy, 0)
	for _, i := range many {
		if i.isHappy() {
			h = append(h, i)
		}
	}
	return h
}

func main() {
	dogs := []dog{
		{"Rocky", 32},
		{"Sam", 63}}

	//Build error: cannot use dogs (type []dog) as type []iHappy in argument to getHappy
	//fmt.Println(getHappy(dogs))

	h := make([]iHappy, len(dogs))
	for i, d := range dogs {
		h[i] = d
	}
	fmt.Println(getHappy(h))
}
