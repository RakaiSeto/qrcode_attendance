package main

import (
	"fmt"
	"time"

	// "os"
	// "strings"

	hook "github.com/robotn/gohook"
)

func main() {
	// add()

	low()
}

func add() {
	fmt.Println("--- Please press ctrl + shift + q to stop hook ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})

	fmt.Println("--- Please press w---")
	hook.Register(hook.KeyDown, []string{"h"}, func(e hook.Event) {
		// theArray := []string{}
	})

	s := hook.Start()
	<-hook.Process(s)
}

var theBool = true
var theTime time.Time

func dor() {
    for ;; {
        if time.Since(theTime) > 2*time.Second{
            if !theBool {
                fmt.Println(theTime)
                fmt.Println("timeout 1")
                theBool = true
            }
        }
    }
}

func low() {
	evChan := hook.Start()
	defer hook.End()
    theArray := []string{}
    go dor()
	for ev := range evChan {
        if ev.Kind == hook.KeyDown {
            theTime = time.Now()
            theArray = append(theArray, string(rune(ev.Keychar)))
            // if strings.Contains(strings.Join(theArray, ""), "http://youtube.com") {
            //     fmt.Println("YOUTUBE")
            //     fmt.Println(strings.Join(theArray, ""))
            //     theArray = nil
            // } else if !strings.HasPrefix(strings.Join(theArray, ""), "h") {
            //     fmt.Println(strings.Join(theArray, ""))
            //     theArray = nil  
            // } else {
            //     fmt.Println(strings.Join(theArray, ""))
            // }
            // fmt.Println(":::", strings.Join(theArray, ""))
            if theBool {
                theBool = false
            }
        }
	}
}