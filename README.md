![go](/img.gif)
# Sort
The sort package sorts data(both string and int) with their number of occurrences and returns the desired data length in descending order with no duplicates of values.
# Use cases range from:
-  ### To implement a trending feature in an e-commerce platform.
-  ### To get the most ordered products e.t.c
# Installation
To install the sort package, you need to first install [Go](https://golang.org/) and set your Go workspace.
1. You can use the below Go command to install sort
```sh
$ go get -u github.com/hisyntax/sort
```
2. Import it in your code:
```sh
import "github.com/hisyntax/sort/x"
```
## Note : Both methods in this package returns one (1) value:
- [x] The sorted array objects in descending order by the length specified.

# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```
## SortInt()
The SortInt() sorts an array of int
```go
package main

import (
	"fmt"

	sorter "github.com/hisyntax/sort/x"
)

func main() {

	arr := []int{2, 3, 10, 4, 1, 10, 1, 4, 4, 5, 6, 6, 6, 6, 1, 6, 7, 8, 12, 9, 1, 1, 1}

    //pass in the array to be sorted and the desired length of data to be returned in descending order
    //the lenght should be zero(0) if you want to get all the sorted data
	sortedInt := sorter.SortInt(arr, 3)
	fmt.Printf("This is the sorted int data: %v\n", sortedInt)

}
```

## SortString()
The SortString() sorts an array of string
```go
package main

import (
	"fmt"

	sorter "github.com/hisyntax/sort/x"
)

func main() {

	arr := []string{"by", "me", "come", "by", "me", "hello", "hey", "hey", "me", "buy", "by", "come", "hello", "go"}

     //pass in the array to be sorted and the desired length of data to be returned in descending order
	sortedString := sorter.SortString(arr, 3)
	fmt.Printf("This is the sorted string data: %v\n", sortedString)

}
```