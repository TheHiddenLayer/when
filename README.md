# `when` is/was it?
The `when` utility takes input as any popular time format and generates a human-readable description of the provided time, indicating whether the time is in the past or future. It also provides a lightweight sense of approximation. This utility obviously enhances any application where timestamped data is frequently displayed to users.

### Usage

```
$ when -v 1443018600
over 7 years ago

$ when 2023-08-13T14:30:00Z
2w

$ when -v 2023-08-13T14:30:00Z
around 2 weeks ago

$ when 2023-09-23T14:30:00Z
09/23/2015

$ when -v -v 1722733189
in roughly 11 months
```