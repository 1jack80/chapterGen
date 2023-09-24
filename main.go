package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// open inFile
	inFile, err := os.Open("./info.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	r := regexp.MustCompile(`(^[\d.\w ?:()|]*).* ([\d]{0,3}:?[\d]{0,3}:[\d]{0,3}).* ([\d]{0,3}:?[\d]{0,3}:[\d]{0,3}).$`)
	outFile, err := os.Create("metadata.txt")
	if err != nil {
		panic("unable to create an output file")
	}

	// for each line
	for scanner.Scan() {
		line := scanner.Text()

		// get title start, end
		parts := r.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}
		title, start, end := parts[1], parts[2], parts[3]

		chapterInfo := fmt.Sprintf(`[CHAPTER] 
TIMEBASE=1/1000 
START= %v
END=%v
title=%v 

`, getms(start), getms(end), title)

		_, err := outFile.WriteString(chapterInfo)
		if err != nil {
			panic(err)
		}

	}
}

func getms(time string) string {
	// fmt.Println(time)
	parts := strings.Split(time, ":")

	sec, min, hrs := "00", "00", "00"
	if len(parts) > 0 {
		sec = parts[len(parts)-1]
		if len(parts) > 1 {
			min = parts[len(parts)-2]
			if len(parts) > 2 {
				hrs = parts[len(parts)-3]
			}
		}
	}

	s, err := strconv.Atoi(sec)
	if err != nil {
		fmt.Printf("time info contains non int character: %v\n ", s)
		panic(nil)
	}
	m, err := strconv.Atoi(min)
	if err != nil {
		fmt.Printf("time info contains non int character: %v\n ", m)
		panic(nil)
	}
	h, err := strconv.Atoi(hrs)
	if err != nil {
		fmt.Printf("time info contains non int character: %v\n ", h)
		panic(nil)
	}

	h = h * 60 * 60 * 1000
	m = m * 60 * 1000
	s = s * 1000

	return fmt.Sprint(h + m + s)
}
