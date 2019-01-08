package main

import (
	"flag"
	"log"

	"github.com/syucream/gitsnap/pkg/gitsnap"
)

func main() {
	path := flag.String("path", ".", "/path/to/.gitdir")
	revision := flag.String("revision", "", "git revision")
	pathPrefix := flag.String("path-prefix", "/tmp/gitsnap/dev/", "/path/to/prefix/destination")
	flag.Parse()

	if err := gitsnap.CreateSnapshotFiles(*path, *pathPrefix, *revision); err != nil {
		log.Fatal(err)
	}
}
