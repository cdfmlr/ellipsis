package ellipsis

import "testing"

func stringWithLength(length int) string {
	s := ""
	for i := 0; i < length; i++ {
		s += "a"
	}
	return s
}

func Test_stringWithLength(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "zero",
			args: args{
				length: 0,
			},
			want: "",
		},
		{
			name: "one",
			args: args{
				length: 1,
			},
			want: "a",
		},
		{
			name: "long",
			args: args{
				length: 10,
			},
			want: "aaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringWithLength(tt.args.length); got != tt.want {
				t.Errorf("stringWithLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func benchmarkEllipsisFunc(b *testing.B, f EllipsisFunc, s string, n int) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f(s, n)
	}
}

func BenchmarkEllipsisCentering0(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(0), 10)
}

func BenchmarkEllipsisCentering1(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(1), 10)
}

func BenchmarkEllipsisCentering10(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(10), 10)
}

func BenchmarkEllipsisCentering30(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(30), 10)
}

func BenchmarkEllipsisCentering100(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(100), 10)
}

func BenchmarkEllipsisCentering1000(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(1000), 10)
}

func BenchmarkEllipsisCentering1000000(b *testing.B) {
	benchmarkEllipsisFunc(b, Centering, stringWithLength(1000000), 10)
}

func BenchmarkEllipsisStarting0(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(0), 10)
}

func BenchmarkEllipsisStarting1(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(1), 10)
}

func BenchmarkEllipsisStarting10(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(10), 10)
}

func BenchmarkEllipsisStarting30(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(30), 10)
}

func BenchmarkEllipsisStarting100(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(100), 10)
}

func BenchmarkEllipsisStarting1000(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(1000), 10)
}

func BenchmarkEllipsisStarting1000000(b *testing.B) {
	benchmarkEllipsisFunc(b, Starting, stringWithLength(1000000), 10)
}

func BenchmarkEllipsisEnding0(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(0), 10)
}

func BenchmarkEllipsisEnding1(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(1), 10)
}

func BenchmarkEllipsisEnding10(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(10), 10)
}

func BenchmarkEllipsisEnding30(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(30), 10)
}

func BenchmarkEllipsisEnding100(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(100), 10)
}

func BenchmarkEllipsisEnding1000(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(1000), 10)
}

func BenchmarkEllipsisEnding1000000(b *testing.B) {
	benchmarkEllipsisFunc(b, Ending, stringWithLength(1000000), 10)
}

func BenchmarkNoEllipsis0(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(0), 10)
}

func BenchmarkNoEllipsis1(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(1), 10)
}

func BenchmarkNoEllipsis10(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(10), 10)
}

func BenchmarkNoEllipsis30(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(30), 10)
}

func BenchmarkNoEllipsis100(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(100), 10)
}

func BenchmarkNoEllipsis1000(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(1000), 10)
}

func BenchmarkNoEllipsis1000000(b *testing.B) {
	benchmarkEllipsisFunc(b, noEllipsis, stringWithLength(1000000), 10)
}
