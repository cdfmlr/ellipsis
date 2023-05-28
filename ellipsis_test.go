// Package ellipsis provides functions to ellipsis a string.
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Centering(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Centering() = %v, want %v", got, tt.want)
			}
		})
	}
}

func FuzzCentering(f *testing.F) {
	for _, s := range []string{
		"",
		"123",
		"1234567890abcdef",
	} {
		for _, n := range []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		} {
			f.Add(s, n)
		}
	}

	f.Fuzz(func(t *testing.T, s string, n int) {
		got := Centering(s, n)

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
	})
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ending(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("Ending() = %v, want %v", got, tt.want)
			}
		})
	}
}
