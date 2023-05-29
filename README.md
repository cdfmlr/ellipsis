# Truncate UTF-8 string with ellipsis

Ellipsis is a Go package that provides functions to truncate long string with ellipsis while being aware of UTF-8 characters, but not spaces. This package offers three main functions to ellipsis a long string: Centering, Starting, and Ending. The functions are aware of UTF-8 characters and will not cut them in half.

## Installation

To use Ellipsis, you need to install Go and set your go mod first. After that, you can get the package by executing the following command:

```bash
go get github.com/cdfmlr/ellipsis
```

## Usage

The package provides three functions that you can use to ellipsis a long string:

### Centering

Centering ellipsis a long string s -> "front...end".

```go
func Centering(s string, n int) string
```

This function takes two arguments: the long string `s` and the maximum length `n` of the ellipsed string. It returns the ellipsed string with the middle part of `s` replaced by "..." to fit into the maximum length.

### Starting

Starting ellipsis a long string s -> "...end".

```go
func Starting(s string, n int) string
```

This function takes two arguments: the long string `s` and the maximum length `n` of the ellipsed string. It returns the ellipsed string with the beginning of `s` replaced by "..." to fit into the maximum length.

### Ending

Ending ellipsis a long string s -> "front...".

```go
func Ending(s string, n int) string
```

This function takes two arguments: the long string `s` and the maximum length `n` of the ellipsed string. It returns the ellipsed string with the end of `s` replaced by "..." to fit into the maximum length.

## Examples

Here are some examples of how to use Ellipsis:

```go
package main

import (
	"fmt"
	"github.com/cdfmlr/ellipsis"
)

func main() {
	s := "0123456789零一二三四五六七八九"

	fmt.Println(Centering(s, 7)) //01...八九

	fmt.Println(Starting(s, 7)) // ...六七八九

	fmt.Println(Ending(s, 7)) // 0123...
}
```

## Contributing

If you find any issues with Ellipsis or have any feature requests, please feel free to submit an issue or a pull request on the project's GitHub page.

## License

Ellipsis is released under the MIT License. See the LICENSE file for details.

## Benchmark

```go
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -coverprofile=/var/folders/dt/b_yjx19j56lb07m0hnmx1zz80000gn/T/vscode-goibiU0i/go-code-cover -bench . github.com/cdfmlr/ellipsis

goos: darwin
goarch: arm64
pkg: github.com/cdfmlr/ellipsis
BenchmarkEllipsisCentering0-8         	523758446	         1.953 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering1-8         	615553364	         1.953 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering10-8        	613988056	         1.956 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering30-8        	12959032	        91.97 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisCentering100-8       	13066100	        91.83 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisCentering1000-8      	13077528	        91.95 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisCentering1000000-8   	13031691	        92.03 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisStarting0-8          	603062298	         1.990 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting1-8          	603817888	         1.988 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting10-8         	602483865	         1.991 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting30-8         	13239217	        89.99 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisStarting100-8        	13304156	        89.94 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisStarting1000-8       	13297048	        90.06 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisStarting1000000-8    	13316767	        90.25 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisEnding0-8            	603063307	         1.990 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding1-8            	602752317	         1.988 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding10-8           	604470055	         1.990 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding30-8           	13481328	        88.56 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisEnding100-8          	13537472	        88.64 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisEnding1000-8         	13493461	        88.47 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisEnding1000000-8      	13489808	        88.49 ns/op	      24 B/op	       2 allocs/op
BenchmarkNoEllipsis0-8                	959663480	         1.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1-8                	948884763	         1.264 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis10-8               	957940106	         1.254 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis30-8               	953301152	         1.260 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis100-8              	952625724	         1.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000-8             	947465343	         1.261 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000000-8          	955488183	         1.265 ns/op	       0 B/op	       0 allocs/op
PASS
	github.com/cdfmlr/ellipsis	coverage: 87.0% of statements
ok  	github.com/cdfmlr/ellipsis	637.262s
```
