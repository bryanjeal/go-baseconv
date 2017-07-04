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

import "testing"

var (
	b62Thousand   = Base62.Encode(1e3)
	b62Million    = Base62.Encode(1e6)
	b6210Million  = Base62.Encode(1e7)
	b62100Million = Base62.Encode(1e8)
)

func TestBase62EncodeDecode(t *testing.T) {
	for i := uint64(0); i < 1e6; i++ {
		got := Base62.Decode(Base62.Encode(i))
		if got != i {
			t.Fatalf("exp %d got %d", i, got)
		}
	}
	for i := ^uint64(0) - 10; i < ^uint64(0); i++ {
		got := Base62.Decode(Base62.Encode(i))
		if got != i {
			t.Fatalf("exp %d got %d", i, got)
		}
	}
}

func BenchmarkBase62EncodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Encode(1e3)
	}
}

func BenchmarkBase62EncodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Encode(1e6)
	}
}

func BenchmarkBase62Encode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Encode(1e7)
	}
}

func BenchmarkBase62Encode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Encode(1e8)
	}
}

func BenchmarkBase62DecodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Decode(b62Thousand)
	}
}

func BenchmarkBase62DecodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Decode(b62Million)
	}
}

func BenchmarkBase62Decode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Decode(b6210Million)
	}
}

func BenchmarkBase62Decode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Base62.Decode(b62100Million)
	}
}
