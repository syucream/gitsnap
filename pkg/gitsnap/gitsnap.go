package gitsnap

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
)

func listFiles(path string, revision string) (*object.FileIter, error) {
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

func copyFile(file *object.File, pathPrefix string) error {
	reader, err := file.Reader()
	if err != nil {
		return err
	}

	writePath := pathPrefix + file.Name
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
}

// CreateSnapshotFiles creates files have content indicated by the git revision.
func CreateSnapshotFiles(gitPath string, destPathPrefix string, revision string) error {
	files, err := listFiles(gitPath, revision)
	if err != nil {
		return err
	}

	return files.ForEach(func(file *object.File) error {
		return copyFile(file, destPathPrefix)
	})
}
