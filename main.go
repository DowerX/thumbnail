package main

import "path/filepath"

func main() {
	conf.ParseConfig()
	path, err := filepath.Abs(conf.Source)
	if err != nil {
		panic(err)
	}
	ProcessDirecotry(conf.Source)
	if conf.Watch {
		InitWatcher()
		SetupWatchers(path)
		RunWatcher()
	}
}
