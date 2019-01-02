# gitsnap

Copy snapshot files managed by git

## Installation

```
$ go get -u github.com/syucream/gitsnap
```

## Usage

```
$ gitsnap -h
Usage of gitsnap:
  -path string
        /path/to/.gitdir (default ".")
  -path-prefix string
        /path/to/prefix/destination (default "/tmp/gitsnap/dev/")
  -revision string
        git revision
```

## Example

```
$ gitsnap -path . -path-prefix /tmp/gitsnap/dev/ -revision master
Done : /tmp/gitsnap/dev/.gitignore
Done : /tmp/gitsnap/dev/Gopkg.lock
Done : /tmp/gitsnap/dev/Gopkg.toml
Done : /tmp/gitsnap/dev/LICENSE
Done : /tmp/gitsnap/dev/README.md
Done : /tmp/gitsnap/dev/main.go
```
