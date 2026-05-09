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


### Day 6:

- [][]bool is clean for on/off grids but [][]int (or a struct with State int) is more flexible when part 2 changes the rules
- max(value, 0) is the clean way to clamp to a minimum in Go 1.21+
- Keep grids separate per part — passing a modified grid to part 2 is a subtle bug that's easy to miss