package main

import (
	"fmt"

	"github.com/ymotongpoo/datemaki"
)

func main() {
	exps := []string{
		"3 days ago",
		"2015 Dec 22nd 23:00:00",
		"yesterday 14:00",
	}

	for _, e := range exps {
		t, err := datemaki.Parse(e)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(t)
	}
}
