# `when`

The `when` utility takes input from any popular time format and generates a human-readable description of the provided time, indicating whether the time is in the past or future. It also provides a lightweight sense of approximation. This utility obviously enhances any application where timestamped data is frequently displayed to users.

### Usage

```bash
$ when 2023-08-13T14:30:00Z
2w

$ when -v 2023-08-13T14:30:00Z
around 2 weeks ago
```