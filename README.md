## `when` ?
The `when` utility takes input as any popular time format and generates a human-readable description of the provided time, indicating whether the time is in the past or future. It also provides a lightweight sense of approximation. This utility obviously enhances any application where timestamped data is frequently displayed to users.

# Installation

```bash
git clone https://github.com/thehiddenlayer/when
make install
```


## Go Usage

```bash
go get github.com/thehiddenlayer/when
```

```go
import "github.com/thehiddenlayer/when"
readableTime, err := when.When(input)
```

Reference `cmd/when/main.go` for more usage in Go.


## CLI Usage

Simply, `when [-v] <timestamp>`

```
$ go install github.com/thehiddenlayer/when/cmd/when@latest
$ when 2023-09-03T09:10:00Z
26m

$ when -v 2023-08-13T14:30:00Z
around 2 weeks ago

$ when -v 1443018600
over 7 years ago

$ when 2023-09-23T14:30:00Z
09/23/2015

$ when -v 1722733189
in roughly 11 months

$ specificDate = 1795130423
$ when -v $someDate
in over 3 years

$ when $specificDate
on 11/19/2026
```
