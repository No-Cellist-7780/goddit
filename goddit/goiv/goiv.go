package goiv

import (
	"fmt"
	"os"
	"runtime"
)

func Viewer(images string, width, height int) {

	if runtime.GOOS == "linux" {
		if os.Getenv("DISPLAY") == "" {
			err := displayDRM(images)
			if err != nil {
				e := displayFB(images)
				if e != nil {
					fmt.Fprintf(os.Stderr, "%s; %s\n", err.Error(), e.Error())
				}
			}
		} else {
			displayX11(images, width, height)
		}
	} else {
		fmt.Println("OS not Supported")
	}
}
