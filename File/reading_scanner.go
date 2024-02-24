package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner
	scanner := bufio.NewScanner(file)

	// scanning for next token in this case \n which is the line delimiter.
	success := scanner.Scan() //read a line
	if success == false {
		// false on error or EOF. Check for errors
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached End Of File.")
		} else {
			log.Fatal(err)
		}
	}

	// Getting the data from the scanner with Bytes() or Text()
	fmt.Println("First Line found:", scanner.Text())
	//If we want the next token, so the next line or \n, we call scanner.Scan() again

	// Reading the whole remaining part of the file:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// Checking for any possible errors:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
