package main

import (
	"context"
	"fmt"
	"password-cracker/cracker"
	"password-cracker/hasher"
	"password-cracker/input"
	"sync"
	"time"
)

const (
	Target       = "$2y$12$hpi41ogGnYphmqiYoGSkZuZso2Azw6fPbRays6lt.KMuYifL4lt0q"
	Routines     = 100
	PasswordList = "./SecLists/Passwords/Common-Credentials/10-million-password-list-top-100.txt"
)

func main() {
	fileInput := input.NewFileInput(PasswordList)
	passwords, err := fileInput.Read()
	if err != nil {
		panic(err)
	}

	getCracking(hasher.NewBCrypt(), passwords)
}

func getCracking(hasher hasher.Hasher, passwords []string) {

	start := makeTimestamp()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := &sync.WaitGroup{}

	max := len(passwords) / Routines
	for i := 0; i < len(passwords); i += max {
		batch := passwords[i:min(i+max, len(passwords))]

		c := cracker.NewCracker(hasher, Target, batch)
		wg.Add(1)
		go c.Crack(wg, ctx)
	}

	wg.Wait()

	end := makeTimestamp()
	fmt.Printf("Done, took %v milliseconds.\n", end-start)
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
