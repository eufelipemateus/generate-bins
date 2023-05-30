package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var csv string

	for i := 0; i < 1000; i++ {
		var bin string
		for i := 0; i < 8; i++ {
			bin = fmt.Sprintf("%s%d", bin, rand.Intn(10))
		}
		csv = fmt.Sprintf("%s%s\n", csv, bin)
	}

	file, err := os.Create("bins.csv")
	println(csv)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(csv)

	if err != nil {
		panic(fmt.Sprintf("failed writing to file: %s", err))
	}

}
