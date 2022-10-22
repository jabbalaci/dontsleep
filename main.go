package main

import (
	"log"
	"os"
	"time"
)

const (
	fname    = "d:/tmp/dontsleep.txt"
	content  = ""
	duration = 5 * time.Minute
	version  = "0.1.0"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	for {
		f, err := os.Create(fname)
		check(err)
		_, err = f.WriteString(content)
		check(err)
		f.Close()
		log.Printf("file %v was created\n", fname)

		time.Sleep(duration)
	}
}
