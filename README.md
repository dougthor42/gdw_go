# gdw_go

A port of my [gdw](https://github.com/dougthor42/gdw) project from
Python to Go.

This is mainly for me to get a feel for Go and see if I like it or not.


## Helpful Links:

+ [Calling Go from Python (StackOverflow)](https://stackoverflow.com/a/56596100/1354930)
+ [Calling Go from Python](https://savorywatt.com/2015/09/18/calling-go-code-from-python-code/)


## Getting Started.


### Installing Go

```shell
# Remove any existing version of Go
rm -rf /usr/local/go
# Download the archive and extract it, creating a tree at /usr/local/go
wget -c https://golang.org/dl/go1.16.2.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
# Add Go to $PATH
# This has already been done so we don't need to do it again.
#export PATH=$PATH:/usr/local/go/bin
# Verify
go version
```


### Set up the project structure.

What we're interested in here is 2-fold:

1.  A Shared Library that can be called from languages such as Python
2.  A CLI client that utilizes the shared library.

```
$ tree
.
├── cmd                         # application entry points. Dir name matches binary
│   └── gdw
│       ├── main.go
│       └── main_test.go
├── LICENSE
├── pkg                         # public consumable code
│   │                           # Seems like the consensus is that there are
│   │                           # no raw .go files in this pkg folder, and
│   │                           # instead everything's put into subfolders.
│   │
│   └── gdw                 # Subfolder for a package
│       │
│       ├── gdw.go          # At least one file, with the same name.
│       │                       # Is this a requirement?
│       └── gdw_test.go
└── README.md
```


### Create the main module

https://golang.org/doc/tutorial/getting-started

The module path *must* be a location from which Go tools can download your
module.

This creates a `go.mod` file in the current directory

```
go mod init github.com/dougthor42/gdw_go
```

If you already have a `.go` file with `package main`, then you can now run

```
go run cmd/gdw/main.go
```


### Build the shared library

See http://snowsyn.net/2016/09/11/creating-shared-libraries-in-go/


### Tests

https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package

It seems like the general consensus is that test files go next to their
counterparts.




## Other notes, unsorted.

It seems like, to run a package (something in the pkg dir), you need to CD
into that dir. You can't just say `go test pkg/gdw/gdw_test.go`.

https://blog.golang.org/using-go-modules

When importing stuff from `pkg`, do this:

```
import "<module>/pkg/gdw"
gdw.Hello()
```

Where `<module>` is what you put when you ran `go mod init`

Functions names are MixedCaps. Variable names are mixedCaps.

https://medium.com/@pliutau/table-driven-tests-in-go-5d7e230681da
