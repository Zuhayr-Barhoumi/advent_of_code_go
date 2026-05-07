### Day 2:

- bufio.Scanner is the idiomatic way to read a file line by line
- strings.Split splits a string by a delimiter into a slice
- strconv.Atoi converts a string to an integer
- defer runs a statement when the surrounding function returns — used to close files
- Built-in min/max accept multiple arguments in Go 1.21+

### Day 4:

- crypto/md5 is built in — md5.Sum([]byte(s)) returns [16]byte
- hex.EncodeToString(bytes[:]) converts to readable hex string
- strings.HasPrefix checks the start of a string
- Bare for {} is Go's infinite loop — use return or break to exit
