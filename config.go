package main

import (
	"github.com/namsral/flag"
)

type Config struct {
	Source      string
	Destination string
	Size        uint
	Force       bool
	Watch       bool
}

var conf Config

func (c *Config) ParseConfig() {
	flag.StringVar(&c.Source, "src", "./", "source")
	flag.StringVar(&c.Destination, "dst", "./.thumbnails", "destination")
	flag.UintVar(&c.Size, "s", 64, "height")
	flag.BoolVar(&c.Force, "force", false, "force overwrite")
	flag.BoolVar(&c.Watch, "watch", false, "watch source for new files")
	flag.Parse()

}
