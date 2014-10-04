package main

import (
	"fmt"

	"github.com/BurntSushi/xgb"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		fmt.Println(err)
		return
	}
    defer X.Close()

}
