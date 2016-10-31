package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    for _,x := range strings.Fields(s){
	    m[x] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

