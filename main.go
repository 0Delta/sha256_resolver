package main

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("please set target hash")
		return
	}
	tgt, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Println(err)
		fmt.Println("invalid hash")
		return
	}
	fmt.Printf("Target : %s\n", os.Args[1])

	chk := func(d []byte) bool {
		dh := sha256.Sum256(d)
		if bytes.Equal(tgt, dh[:]) {
			return true
		}
		return false
	}

	ans := make(chan string, 1)
	lim := make(chan bool, 16)
	c := []byte{0}
	ctx, cancel := context.WithCancel(context.Background())
	end := false

	for !end {
		lim <- true
		select {
		case <-ctx.Done():
			end = true
			break
		default:
			cc := make([]byte, len(c))
			copy(cc, c)
			go func(c []byte) {
				defer func() { <-lim }()
				for i := 0; i < 0xFF*0xFF; i++ {
					if chk(c) {
						ans <- string(c)
						cancel()
					}
					c = byte_increment(c, 0)
				}
			}(cc)
			if len(c) < 2 {
				c = []byte{0, 0, 0}
			} else {
				c = byte_increment(c, 2)
			}
			continue
		}
	}
	anser := <-ans
	fmt.Printf("anser : %s\n", anser)
	close(ans)
	close(lim)
}

func byte_increment(b []byte, idx int) []byte {
	b[idx] += 1
	if b[idx] == 0x00 {
		if len(b) <= idx+1 {
			b = append([]byte{0x01}, b...)
		} else {
			return byte_increment(b, idx+1)
		}
	}
	return b
}
