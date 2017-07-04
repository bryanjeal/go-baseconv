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

package baseconv

type baseNAEncoding struct {
	encode    [47]byte
	decodeMap [256]uint64
}

// BaseNoAmbiguous namespaces the baseNA encoder
var BaseNoAmbiguous = &baseNAEncoding{
	[47]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'm', 'n',
		'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'C',
		'E', 'F', 'H', 'J', 'K', 'L', 'M', 'N', 'P', 'R', 'T', 'U', 'V', 'W',
		'X', 'Y', '3', '4', '7', '9'},

	[256]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 43, 44, 0, 0, 45, 0, 46, 0, 0, 0, 0, 0, 0, 0, 25, 0,
		26, 0, 27, 28, 0, 29, 0, 30, 31, 32, 33, 34, 0, 35, 0, 36, 0, 37, 38,
		39, 40, 41, 42, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		0, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0},
}

// Decode takes a number that is encoded in baseNA and returns it as an uint64
func (e *baseNAEncoding) Decode(buf []byte) uint64 {
	sum := uint64(0)

	for i := 0; i < len(buf); i++ {
		c := buf[i]
		sum = sum*uint64(47) + e.decodeMap[c]
	}

	return sum
}

// Encode takes an uint64 and returns it encoded as a baseNA []byte
func (e *baseNAEncoding) Encode(n uint64) []byte {
	if n == 0 {
		return []byte{e.encode[0]}
	}

	buf := make([]byte, 12)
	i := 12

	for n > 0 && i >= 0 {
		i--
		buf[i] = e.encode[n%47]
		n /= 47
	}

	return buf[i:]
}
