// build a command line tool that functions like unix wc command
// wc -l <filename>
// wc -w <filename>
// wc -c <filename>
// wc -m <filename>
// wc -L <filename>
// wc -help
// wc -version

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// get the command line
	flag.Usage = func() {
		fmt.Println("Usage: wc [OPTION]... [FILE]...")
		fmt.Println("Print newline, word, and byte counts for each FILE, and a total line if more than one FILE is specified.  With no FILE, or when FILE is -, read standard input.")
		fmt.Println("  -c, --bytes            print the byte counts")
		fmt.Println("  -m, --chars            print the character counts")
		fmt.Println("  -l, --lines            print the newline counts")
		fmt.Println("  -L, --max-line-length  print the length of the longest line")
		fmt.Println("  -w, --words            print the word counts")
		fmt.Println("      --help     display this help and exit")
		fmt.Println("      --version  output version information and exit")
	}

	countBytes := flag.Bool("c", false, "print the byte counts")
	countChars := flag.Bool("m", false, "print the character counts")
	countLines := flag.Bool("l", false, "print the newline counts")
	countMaxLineLength := flag.Bool("L", false, "print the length of the longest line")
	countWords := flag.Bool("w", false, "print the word counts")
	flag.Bool("help", false, "display this help and exit")
	flag.Bool("version", false, "output version information and exit")
	// accept combinations of flags

	// parse the flags
	flag.Parse()
	// get the filename
	filename := flag.Arg(0)
	// check if the file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return
	}
	
	// check if the file exists
	fmt.Println("File exists")
	fmt.Printf("Reading file %s\n", filename)
	// read the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()
	// create reader
	reader := bufio.NewReader(file)
	
	counts := make(map[string]int)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
		 break
		}

		if(*countBytes){
			counts["bytes"] += len(line)
		}
		if(*countChars){
			counts["chars"] += len(line)
		}
		if(*countLines){
			counts["lines"]++
		}
		if(*countMaxLineLength){
			if len(line) > counts["maxLineLength"] {
				counts["maxLineLength"] = len(line)
			}
		}
		if(*countWords){
			words := split(line)
			counts["words"] += len(words)
		}
	}

flag.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "c":
			fmt.Printf("Byte count: %d\n", counts["bytes"])
		case "m":
			fmt.Printf("Character count: %d\n", counts["chars"])
		case "l":
			fmt.Printf("Line count: %d\n", counts["lines"])
		case "L":
			fmt.Printf("Max line length: %d\n", counts["maxLineLength"])
		case "w":
			fmt.Printf("Word count: %d\n", counts["words"])
		default:
			fmt.Printf("Character count: %d\n", counts["chars"])
			fmt.Printf("Word count: %d\n", counts["words"])
			fmt.Printf("Line count: %d\n", counts["lines"])
		}
	})


	

}

func split(line string) []string {
	return strings.Fields(line)
}


