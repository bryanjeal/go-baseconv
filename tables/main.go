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

const digits = "0123456789"
const uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowers = "abcdefghijklmnopqrstuvwxyz"
const ambiguous = "B8G6I1l0OQDS5Z2"

func main() {
	// No Ambiguous Characters
	fmt.Println("---------------------------No Ambiguous Characters----------------------------")
	baseNoAmbigous()
	fmt.Println("-------------------------------------------------------------------------------")

	// Base62
	fmt.Println("\n------------------------------base62 (a-zA-Z0-9)------------------------------")
	base62()
	fmt.Println("-------------------------------------------------------------------------------")
}
