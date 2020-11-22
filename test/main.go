package main

import (
	"github.com/chenguofan1999/cloudgo"
)

func main() {
	n := cloudgo.NewServer()
	n.Run(":5990")
}
