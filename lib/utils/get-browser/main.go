// Package main ...
package main

import (
	"fmt"

	"github.com/yeyu12/rod/lib/launcher"
	"github.com/yeyu12/rod/lib/utils"
)

func main() {
	p, err := launcher.NewBrowser().Get()
	utils.E(err)

	fmt.Println(p)
}
