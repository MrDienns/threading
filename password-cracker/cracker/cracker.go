package cracker

import (
	"context"
	"fmt"
	"password-cracker/hasher"
	"sync"
)

type Cracker struct {
	hasher    hasher.Hasher
	target    string
	passwords []string
}

func NewCracker(hasher hasher.Hasher, target string, passwords []string) *Cracker {
	return &Cracker{hasher: hasher, target: target, passwords: passwords}
}

func (c *Cracker) Crack(wg *sync.WaitGroup, ctx context.Context) {

	defer wg.Done()

	for _, password := range c.passwords {
		matches := c.hasher.Compare(c.target, password)

		if matches {

			fmt.Print("Found password: " + password + "\n")
			ctx.Done()
			return
		}
	}
}
