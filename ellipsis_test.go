package ellipsis

import "testing"

func TestCentering(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s: "",
				n: 10,
			},
			want: "",
		},
		{
			name: "short_s",
			args: args{
				s: "123",
				n: 10,
			},
			want: "123",
		},
		{
			name: "short_n",
			args: args{
				s: "123",
				n: 1,
			},
			want: "...",
		},
		{
			name: "long_even",
			args: args{
				s: "1234567890abcdef",
				n: 10,
			},
			want: "123...def",
		},
		{
			name: "long_odd",
			args: args{
				s: "1234567890abcdef",
				n: 11,
			},
			want: "1234...cdef",
		},
		{
			name: "utf8",
			args: args{
				s: "一二三四五六七八九十",
				n: 7,
			},
			want: "一二...九十",
		},
		{
			name: "bad_utf8",
			args: args{
				s: "\xbc12345\xe3",
				n: 5,
			},
			want: "\uFFFD...\uFFFD", // \uFFFD is the replacement character ('�') a.k.a. RuneError
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Centering(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Centering() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStarting(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s: "",
				n: 10,
			},
			want: "",
		},
		{
			name: "short_s",
			args: args{
				s: "123",
				n: 10,
			},
			want: "123",
		},
		{
			name: "short_n",
			args: args{
				s: "123",
				n: 1,
			},
			want: "...",
		},
		{
			name: "long_even",
			args: args{
				s: "1234567890abcdef",
				n: 10,
			},
			want: "...0abcdef",
		},
		{
			name: "long_odd",
			args: args{
				s: "1234567890abcdef",
				n: 11,
			},
			want: "...90abcdef",
		},
		{
			name: "utf8",
			args: args{
				s: "一二三四五六七八九十",
				n: 7,
			},
			want: "...七八九十",
		},
		{
			name: "bad_utf8",
			args: args{
				s: "\xbc12345\xe3",
				n: 5,
			},
			want: "...5\uFFFD", // \uFFFD is the replacement character ('�') a.k.a. RuneError
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Starting(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Starting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnding(t *testing.T) {
	type args struct {
		s string
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				s: "",
				n: 10,
			},
			want: "",
		},
		{
			name: "short_s",
			args: args{
				s: "123",
				n: 10,
			},
			want: "123",
		},
		{
			name: "short_n",
			args: args{
				s: "123",
				n: 1,
			},
			want: "...",
		},
		{
			name: "long_even",
			args: args{
				s: "1234567890abcdef",
				n: 10,
			},
			want: "1234567...",
		},
		{
			name: "long_odd",
			args: args{
				s: "1234567890abcdef",
				n: 11,
			},
			want: "12345678...",
		},
		{
			name: "utf8",
			args: args{
				s: "一二三四五六七八九十",
				n: 7,
			},
			want: "一二三四...",
		},
		{
			name: "bad_utf8",
			args: args{
				s: "\xbc12345\xe3",
				n: 5,
			},
			want: "\uFFFD1...", // \uFFFD is the replacement character ('�') a.k.a. RuneError
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ending(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Ending() = %v, want %v", got, tt.want)
			}
		})
	}
}

// FuzzTest is a fuzz test for all three ellipsis functions.
func FuzzTest(f *testing.F) {
	for _, s := range []string{
		"",
		"123",
		"1234567890abcdef",
		"一二三四五六七八九十",
	} {
		for _, n := range []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		} {
			f.Add(s, n)
		}
	}

	f.Fuzz(func(t *testing.T, s string, n int) {
		for _, ellipsisFunc := range []ellipsisFunc{
			Centering, Starting, Ending,
		} {
			got := ellipsisFunc(s, n)

			if n <= 3 {
				if got != "..." {
					t.Errorf("Centering(%q, %d) = %q, want %q", s, n, got, "...")
				}
				return
			}

			r := []rune(got)

			if len(r) > n {
				t.Errorf("Centering(%q, %d) = %q (len = %v), want len %v", s, n, got, len(r), n)
				return
			}
		}
	})
}
