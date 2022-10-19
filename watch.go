package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func InitWatcher() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
}

func SetupWatchers(path string) {
	fmt.Println("Adding", path)
	if watcher == nil {
		InitWatcher()
	}
	dir, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	watcher.Add(path)
	for _, e := range dir {
		if e.IsDir() {
			SetupWatchers(filepath.Join(path, e.Name()))
		}
	}
}

func RunWatcher() {
	for {
		select {
		case e, ok := <-watcher.Events:
			if !ok {
				return
			}
			if e.Has(fsnotify.Create) {
				root, _ := filepath.Abs(filepath.Join(conf.Destination, ".."))
				path, err := filepath.Rel(root, e.Name)
				path = filepath.Join(conf.Destination, path)
				if err != nil {
					fmt.Println("Event error: path:", err)
				}
				info, _ := os.Stat(e.Name)
				if info.IsDir() {
					fmt.Println("New directory:", path)
					os.MkdirAll(path, os.ModePerm)
					watcher.Add(e.Name)
				} else {
					fmt.Println("Working on:", e.Name, path)
					for i := 1; i < 5; i++ {
						err = nil
						err = ResizeImage(e.Name, path)
						if err == nil {
							break
						}
						time.Sleep(time.Second)
					}
					if err != nil {
						fmt.Println("Event error:", err)
					}
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			if err != nil {
				panic(err)
			}
		}
	}
}
