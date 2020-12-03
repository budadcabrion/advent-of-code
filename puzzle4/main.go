package main
 
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)
 
func main() {
	file, err := os.Open("data.txt")
	defer file.Close()
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
 
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
 
 	total := 0
 	valid := 0

	for _, line := range lines {
		//fmt.Println(line)
		total = total + 1
		if isLineValid(line) {
			valid = valid + 1
		}
	}
	fmt.Println("Valid Lines:", valid, "/", total)
}

func isLineValid(line string) bool {
	line = strings.ReplaceAll(line, "-", " ")
	line = strings.ReplaceAll(line, ": ", " ")
	
	p := strings.Split(line, " ")
	a, _ := strconv.Atoi(p[0])
	b, _ := strconv.Atoi(p[1])
	letter := p[2]
	password := p[3]
	
	aLetter := string(password[a-1])
	bLetter := string(password[b-1])

	aMatch := aLetter == letter
	bMatch := bLetter == letter

	isValid := aMatch != bMatch	

	fmt.Println(a, b, letter, password, "-", aLetter, bLetter, isValid)

	return isValid
}