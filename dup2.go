package main

import ( 
	"bufio" 
	"fmt" 
	"os"
)

func main() { 
	counts := make(map[string]int) 
	files := os.Args[1:] 
	for _, arg := range files { 
		f, err := os.Open(arg) 
		if err != nil { 
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err) 
			continue 
		} 
		if countLines(f, counts) {
			fmt.Printf("%v\n", arg) 
		} 
		f.Close() 
	}  
}

func countLines(f *os.File, counts map[string]int) bool { 
	input := bufio.NewScanner(f) 
	for input.Scan() { 
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			return true
		}
	}

	return false
}