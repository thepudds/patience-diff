# patience-diff

Trivial command line utility:

```
$ go install github.com/thepudds/patience-diff/cmd/patience-diff@latest

$ patience-diff file1 file2
```

The diff implementation is a copy of a fast & memory efficient diff implementation from a Go stdlib internal package — [patience diff](https://bramcohen.livejournal.com/73318.html?noscroll&utm_medium=endless_scroll) as implemented by Russ Cox in [CL 384255](https://golang.org/cl/384255/). The CL has not yet been merged into the main Go repo.
