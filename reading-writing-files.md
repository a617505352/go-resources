<h1 align="center">Reading and Writing Files</h1>

- [Reading Files](#reading-files)
- [Writing Files to New Files](#writing-files-to-new-files)
- [Writing to Existing Files](#writing-to-existing-files)

# Reading Files

```go
package main

// import the 2 modules we need
import (
	"fmt"
	"io/ioutil"
)

func main() {
  // read in the contents of the localfile.data
	data, err := ioutil.ReadFile("localfile.data")
  // if our program was unable to read the file
  // print out the reason why it can't
  if err != nil {
		fmt.Println(err)
	}
  // if it was successful in reading the file then
  // print out the contents as a string
	fmt.Print(string(data))
}
```

[[↑] Back to top](#reading-and-writing-files)

# Writing Files to New Files

```go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	mydata := []byte("All the data I wish to write to a file")
	// the WriteFile method returns an error if unsuccessful
	err := ioutil.WriteFile("myfile.data", mydata, 0777)
	if err != nil {
		fmt.Println(err)
  }

	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}
```

[[↑] Back to top](#reading-and-writing-files)

# Writing to Existing Files

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.OpenFile("myfile.data", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.WriteString("new data that wasn't there originally\n")
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadFile("myfile.data")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(data))
}
```

[[↑] Back to top](#reading-and-writing-files)
