package main

import (
	"fmt"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	X, err := xgb.NewConnDisplay(":0")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer X.Close()

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)
	xproto.ChangeWindowAttributes(X, screen.Root, xproto.CwBackPixel|xproto.CwEventMask, []uint32{0xffffffff, xproto.EventMaskStructureNotify | xproto.EventMaskKeyPress | xproto.EventMaskKeyRelease})

	for {
		ev, err := X.WaitForEvent()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Event: %s", ev)
	}

}
