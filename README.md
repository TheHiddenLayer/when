## `when` is or was it?
The `when` utility takes input as any popular time format and generates a human-readable description of the provided time, indicating whether the time is in the past or future. It also provides a lightweight sense of approximation. This utility obviously enhances any application where timestamped data is frequently displayed to users.

Think of any microblogging platform:

<img height="500" src="https://raw.githubusercontent.com/the-hidden-layer/assets/master/img/F3441A73-8B02-4184-AD2A-8479FBD0C4DA_1_201_a.jpeg" />

# Installation

```bash
git clone https://github.com/the-hidden-layer/when
make install
```

## Go Usage

```bash
go get github.com/the-hidden-layer/when
```

```go
import "github.com/the-hidden-layer/when"
readableTime, err := when.When(input)
```

Reference `cmd/when/main.go` for more usage in Go.


## CLI Usage

Simply, `when [-v] <timestamp>`

```
$ go install github.com/the-hidden-layer/when/cmd/when@latest
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

$ someDate = 1795130423
$ when -v $someDate
in over 3 years

$ when $someDate
on 11/19/2026
```
