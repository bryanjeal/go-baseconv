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
	bNAThousand   = BaseNoAmbiguous.Encode(1e3)
	bNAMillion    = BaseNoAmbiguous.Encode(1e6)
	bNA10Million  = BaseNoAmbiguous.Encode(1e7)
	bNA100Million = BaseNoAmbiguous.Encode(1e8)
)

func TestBaseNoAmbiguousEncodeDecode(t *testing.T) {
	for i := uint64(0); i < 1e6; i++ {
		got := BaseNoAmbiguous.Decode(BaseNoAmbiguous.Encode(i))
		if got != i {
			t.Fatalf("exp %d got %d", i, got)
		}
	}
	for i := ^uint64(0) - 10; i < ^uint64(0); i++ {
		got := BaseNoAmbiguous.Decode(BaseNoAmbiguous.Encode(i))
		if got != i {
			t.Fatalf("exp %d got %d", i, got)
		}
	}
}

func BenchmarkBaseNoAmbiguousEncodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Encode(1e3)
	}
}

func BenchmarkBaseNoAmbiguousEncodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Encode(1e6)
	}
}

func BenchmarkBaseNoAmbiguousEncode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Encode(1e7)
	}
}

func BenchmarkBaseNoAmbiguousEncode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Encode(1e8)
	}
}

func BenchmarkBaseNoAmbiguousDecodeThousand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Decode(bNAThousand)
	}
}

func BenchmarkBaseNoAmbiguousDecodeMillion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Decode(bNAMillion)
	}
}

func BenchmarkBaseNoAmbiguousDecode10Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Decode(bNA10Million)
	}
}

func BenchmarkBaseNoAmbiguousDecode100Million(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BaseNoAmbiguous.Decode(bNA100Million)
	}
}
