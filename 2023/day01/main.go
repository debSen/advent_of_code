package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	fptr := flag.String("fpath", "input.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	count := 0
	for s.Scan() {
		re := regexp.MustCompile("[0-9]")
		line := s.Text()
		pattern_match := re.FindAllString(line, -1)
		res := 0

		i, _ := strconv.Atoi(pattern_match[0])
		if len(pattern_match) == 1 {
			res = i * 11
		} else {
			j, _ := strconv.Atoi(pattern_match[len(pattern_match)-1])
			res = i*10 + j
		}
		count += res
		// fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
