package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("/tmp/file.txt"); err != nil {
		return
	}
	fmt.Printf("%x\n", md5.Sum([]byte("Rem pafon hopcak okmode zumuzap.")))
}
