//Build: go build -o logParser main.go
//Usage: ./logParser text.log NOK

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <filename> state (OK|NOK)", os.Args[0])
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("Can't open file ", os.Args[1], err)
	}
	defer file.Close()

	expr := `^\[(\d{4}-\d{2}-\d{2})\s+(\d{2}:\d{2}):\d{2}\]\s+` + os.Args[2]
	r := regexp.MustCompile(expr)
	scanner := bufio.NewScanner(file)
	result := map[string]int{}
	for scanner.Scan() {
		if r.MatchString(scanner.Text()) {
			key := r.ReplaceAllString(scanner.Text(), "$1 $2")
			result[key]++
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for k, v := range result {
		fmt.Println(k, ":", v)
	}
}
