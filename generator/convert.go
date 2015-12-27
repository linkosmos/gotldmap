package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/linkosmos/gotldmap"
)

func lastMark() (last int) {
	for _, mark := range gotldmap.Map {
		if mark > last {
			last = mark
		}
	}
	return last
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	count := lastMark()
	count += 1 // Increment over last mark
	scanner := bufio.NewScanner(file)

	fmt.Println("package gotldmap")
	fmt.Println("")
	fmt.Println("// Generated; DO NOT EDIT")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("// Map - map of top level domains with mark key")
	fmt.Println("var Map = map[string]int{")
	for scanner.Scan() {
		tld := strings.ToLower(scanner.Text())
		if gotldmap.TldExist(tld) {
			tldPos, _ := gotldmap.FindByTld(tld)
			fmt.Printf("\"%s\": %d,\n", tld, tldPos)
		} else {
			fmt.Printf("\"%s\": %d,\n", tld, count)
			count++
		}

	}
	fmt.Println("}")
	os.Exit(1)
}
