/////////////////////////////////
// Creating, Opening, Closing, Renaming, Moving, and Removing files in Go
// Go Playground: https://play.golang.org/p/Sz_LfNS9GKU
/////////////////////////////////

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// CREATING A FILE
	newFile, err := os.Create("a.txt")
	if err != nil {
		log.Fatal(err)

	}

	// TRUNCATING A FILE
	err = os.Truncate("a.txt", 0) //0 means completely empty the file.
	if err != nil {
		log.Fatal(err)
	}

	// CLOSING THE FILE
	newFile.Close()

	// OPEN AND CLOSE AN EXISTING FILE
	file, err := os.Open("a.txt") // open in read-only mode

	// error handling
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	//OPENING a FILE WITH MORE OPTIONS
	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// GETTING FILE INFO
	var fileInfo os.FileInfo
	fileInfo, err = os.Stat("a.txt")

	p := fmt.Println
	p("File Name:", fileInfo.Name())        // => File Name: a.txt
	p("Size in bytes:", fileInfo.Size())    // => Size in bytes: 0
	p("Last modified:", fileInfo.ModTime()) // => Last modified: 2019-10-21 16:16:00.325037748 +0300 EEST
	p("Is Directory? ", fileInfo.IsDir())   // => Is Directory?  false
	p("Pemissions:", fileInfo.Mode())       // => Pemissions: -rw-r-----

	// CHECKING IF FILE EXISTS
	fileInfo, err = os.Stat("b.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	// RENAMING AND MOVING A FILE
	oldPath := "a.txt"
	newPath := "aaa.txt"
	err = os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	// REMOVING A FILE
	err = os.Remove("aa.txt")
	if err != nil {
		log.Fatal(err)
	}

}
