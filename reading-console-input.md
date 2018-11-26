<h1 align="center">Reading Console Input</h1>

- [Using Bufio’s Scanner](#using-bufios-scanner)

# Using Bufio’s Scanner

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```
