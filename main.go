package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

const (
	path = ""
	revision = ""
	writePathPrefix = "/tmp/gitsnap/dev/"
)

func listFiles(path string, revision string) (*object.FileIter, error){
	repos, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	hash, err := repos.ResolveRevision(plumbing.Revision(revision))
	if err != nil {
		return nil, err
	}

	commit, err := repos.CommitObject(*hash)
	if err != nil {
		return nil, err
	}

	tree, err := commit.Tree()
	if err != nil {
		return nil, err
	}

	return tree.Files(), nil
}

func writeFile(reader io.Reader, path string) error {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, reader)
	return err
}

func main() {
	files, err := listFiles(path, revision)
	if err != nil {
		log.Fatal(err)
	}

	err = files.ForEach(func(file *object.File) error {
		reader, err := file.Reader()
		if err != nil {
			return err
		}

		writePath := writePathPrefix + file.Name
		err = writeFile(reader, writePath)
		if err != nil {
			return err
		}

		err = reader.Close()
		if err != nil {
			return err
		}

		fmt.Println("Done : " + writePath)

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}