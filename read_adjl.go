package main

// read in a graph through its adj-list.

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func check_err(e error) {
	if e != nil {
		panic(e)
	}
}

func exist_file(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func proc_pp(inp string) []int {
	var ppIDs []int
	// ... pp is prev or post :-)
	pp := strings.Split(inp, "~")
	// ...
	if len(pp) < 1 {
		fmt.Println("Error! too few node IDs are provided")
		ppIDs = make([]int, 0)
	} else if len(pp) == 1 {
		// a single node id is provided.
		id, err := strconv.Atoi(strings.TrimSpace(pp[0]))
		check_err(err)
		ppIDs = make([]int, 1)
		ppIDs[0] = id
	} else if len(pp) == 2 {
		// a nodes ids list is provided.
		id0, err := strconv.Atoi(strings.TrimSpace(pp[0]))
		check_err(err)
		id1, err := strconv.Atoi(strings.TrimSpace(pp[1]))
		check_err(err)
		// ...
		if id1 < id0 {
			fmt.Print("Error list! Find a reversed range!")
			ppIDs = make([]int, 0)
		} else {
			ppIDs = make([]int, id1-id0+1)
			for k := id0; k <= id1; k++ {
				ppIDs[k-id0] = k
			}
		}
	} else if len(pp) > 2 {
		fmt.Println("Error! too many node IDs are provided")
		ppIDs = make([]int, 0)
	}
	return ppIDs
}

func main() {
	var fAdjlName string = "a.adjl"
	if exist_file(fAdjlName) {
		data, err := ioutil.ReadFile(fAdjlName)
		check_err(err)
		aaa := strings.Split(string(data), "\n")
		for i := range aaa {
			fmt.Println(i, " : ", aaa[i])
			if len(aaa[i]) <= 0 {
				fmt.Println("--- \t ((this is empty line ( not processed")
			} else if aaa[i][0] == '#' {
				fmt.Println("--- \t ((this is a comment ( not processed")
			} else if aaa[i][0] == '=' {
				fmt.Println("--- \t ((this is reserved for global settings ( not processed")
			} else {
				fmt.Print("--- \t ")
				bbb := strings.Split(aaa[i], ":")
				prevIDs := proc_pp(bbb[0])
				postIDs := proc_pp(bbb[1])
				for x := range prevIDs {
					for y := range postIDs {
						fmt.Print(prevIDs[x], ":", postIDs[y], "\t")
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}
	}
}
