// Copyright 2017 Bryan Jeal <bryan@jeal.ca>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import "fmt"

func printAlphabet(alphabet []byte) {
	fmt.Printf("base: %d\n", len(alphabet))
	fmt.Printf("encode: [base]byte{")
	for i, c := range alphabet {
		fmt.Printf("'%c'", c)
		if i != len(alphabet)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")

	max := (1 << 8)
	hash := make([]int, max)

	for i := 1; i < max; i++ {
		for j := 1; j < len(alphabet); j++ {
			if i == int(alphabet[j]) {
				hash[i] = j
				break
			}
		}
	}

	fmt.Printf("\ndecodeMap: [256]uint64{")

	for i, d := range hash {
		fmt.Printf("%d", d)
		if i != len(hash)-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("}\n")
}
