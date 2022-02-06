package main;

import (
	"fmt"

	"github.com/zhuangjinjin/gtools"
)

func main() {
	result := gtools.VerifyIdCard("350521XXXXXXXXXXX8")
	fmt.Println(result)
}