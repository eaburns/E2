package main

import (
	"log"
	"os"
	"path"

	"9fans.net/go/acme"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: E2 <file>")
	}

	win, err := acme.New()
	if err != nil {
		log.Fatal(err)
	}

	win.Name(name())
	win.Write("tag", []byte("Put "))
	win.Ctl("get")

	for {
		select {
		case ev, ok := <-win.EventChan():
			if !ok {
				return
			}
			win.WriteEvent(ev)

			if (ev.C2 == 'x' || ev.C2 == 'X') && string(ev.Text) == "Put" {
				win.Del(true)
			}
		}
	}
}

func name() string {
	name := os.Args[1]
	if !path.IsAbs(name) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		name = path.Join(wd, name)
	}
	return path.Clean(name)
}
