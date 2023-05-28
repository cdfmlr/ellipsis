# ellipsis

## Benchmark

`r := []rune(s)`:

```go
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -coverprofile=/var/folders/dt/b_yjx19j56lb07m0hnmx1zz80000gn/T/vscode-goibiU0i/go-code-cover -bench . github.com/cdfmlr/ellipsis

goos: darwin
goarch: arm64
pkg: github.com/cdfmlr/ellipsis
BenchmarkEllipsisCentering0-8         	189445855	         5.942 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering1-8         	198618291	         6.019 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering10-8        	86998551	        13.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering30-8        	13403720	        88.78 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisCentering100-8       	 6237844	       189.0 ns/op	     440 B/op	       3 allocs/op
BenchmarkEllipsisCentering1000-8      	 1000000	      1187 ns/op	    4120 B/op	       3 allocs/op
BenchmarkEllipsisCentering1000000-8   	    1197	   1011198 ns/op	 4005917 B/op	       3 allocs/op
BenchmarkEllipsisStarting0-8          	231353304	         5.175 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting1-8          	218346172	         5.498 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting10-8         	85884010	        13.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting30-8         	13407708	        89.06 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisStarting100-8        	 6291519	       189.1 ns/op	     440 B/op	       3 allocs/op
BenchmarkEllipsisStarting1000-8       	 1000000	      1187 ns/op	    4120 B/op	       3 allocs/op
BenchmarkEllipsisStarting1000000-8    	    1141	   1025078 ns/op	 4005917 B/op	       3 allocs/op
BenchmarkEllipsisEnding0-8            	194571955	         6.140 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding1-8            	192806509	         6.212 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding10-8           	85051340	        14.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding30-8           	13416646	        89.17 ns/op	      24 B/op	       2 allocs/op
BenchmarkEllipsisEnding100-8          	 6352474	       187.5 ns/op	     440 B/op	       3 allocs/op
BenchmarkEllipsisEnding1000-8         	 1000000	      1186 ns/op	    4120 B/op	       3 allocs/op
BenchmarkEllipsisEnding1000000-8      	    1168	   1030456 ns/op	 4005916 B/op	       3 allocs/op
BenchmarkNoEllipsis0-8                	957221186	         1.255 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1-8                	949984315	         1.261 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis10-8               	957951255	         1.266 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis30-8               	946612692	         1.283 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis100-8              	951024541	         1.258 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000-8             	949726180	         1.256 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000000-8          	944248982	         1.260 ns/op	       0 B/op	       0 allocs/op
PASS
	github.com/cdfmlr/ellipsis	coverage: 90.9% of statements
ok  	github.com/cdfmlr/ellipsis	472.310s
```

`cutString`:

```go
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -coverprofile=/var/folders/dt/b_yjx19j56lb07m0hnmx1zz80000gn/T/vscode-goibiU0i/go-code-cover -bench . github.com/cdfmlr/ellipsis

goos: darwin
goarch: arm64
pkg: github.com/cdfmlr/ellipsis
BenchmarkEllipsisCentering0-8         	542044388	         1.922 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering1-8         	617481146	         1.974 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering10-8        	609189238	         1.968 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering30-8        	 6926486	       169.7 ns/op	      40 B/op	       8 allocs/op
BenchmarkEllipsisCentering100-8       	 7061580	       170.1 ns/op	      40 B/op	       8 allocs/op
BenchmarkEllipsisCentering1000-8      	 7070059	       169.8 ns/op	      40 B/op	       8 allocs/op
BenchmarkEllipsisCentering1000000-8   	 7062255	       169.6 ns/op	      40 B/op	       8 allocs/op
BenchmarkEllipsisStarting0-8          	625284145	         1.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting1-8          	625439761	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting10-8         	627421454	         1.918 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting30-8         	 6396756	       186.6 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisStarting100-8        	 6423033	       186.7 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisStarting1000-8       	 6428372	       186.7 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisStarting1000000-8    	 6387790	       187.5 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisEnding0-8            	626359966	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding1-8            	625971151	         1.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding10-8           	625293649	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding30-8           	 6307149	       189.3 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisEnding100-8          	 6322581	       189.1 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisEnding1000-8         	 6355480	       188.8 ns/op	      64 B/op	       9 allocs/op
BenchmarkEllipsisEnding1000000-8      	 6334530	       189.5 ns/op	      64 B/op	       9 allocs/op
BenchmarkNoEllipsis0-8                	948885700	         1.259 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1-8                	940724553	         1.262 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis10-8               	954214596	         1.260 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis30-8               	941029167	         1.261 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis100-8              	947669245	         1.254 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000-8             	950792524	         1.263 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000000-8          	956940660	         1.260 ns/op	       0 B/op	       0 allocs/op
PASS
	github.com/cdfmlr/ellipsis	coverage: 89.4% of statements
ok  	github.com/cdfmlr/ellipsis	644.004s

```

`runes[i] = char`:

```go
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -coverprofile=/var/folders/dt/b_yjx19j56lb07m0hnmx1zz80000gn/T/vscode-goibiU0i/go-code-cover -bench . github.com/cdfmlr/ellipsis

goos: darwin
goarch: arm64
pkg: github.com/cdfmlr/ellipsis
BenchmarkEllipsisCentering0-8         	540390734	         1.916 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering1-8         	621857242	         1.932 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering10-8        	627413390	         1.914 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisCentering30-8        	10521423	       117.0 ns/op	      72 B/op	       6 allocs/op
BenchmarkEllipsisCentering100-8       	 9447420	       123.8 ns/op	      72 B/op	       6 allocs/op
BenchmarkEllipsisCentering1000-8      	10249033	       113.7 ns/op	      72 B/op	       6 allocs/op
BenchmarkEllipsisCentering1000000-8   	10483939	       114.0 ns/op	      72 B/op	       6 allocs/op
BenchmarkEllipsisStarting0-8          	626774361	         1.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting1-8          	626343758	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting10-8         	627081016	         1.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisStarting30-8         	11962456	       100.4 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisStarting100-8        	11953460	       100.5 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisStarting1000-8       	11955484	       100.4 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisStarting1000000-8    	11798694	       101.2 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisEnding0-8            	625710849	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding1-8            	626564917	         1.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding10-8           	626310387	         1.917 ns/op	       0 B/op	       0 allocs/op
BenchmarkEllipsisEnding30-8           	11714025	       102.1 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisEnding100-8          	11752506	       102.1 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisEnding1000-8         	11753288	       102.1 ns/op	      72 B/op	       4 allocs/op
BenchmarkEllipsisEnding1000000-8      	11672947	       102.6 ns/op	      72 B/op	       4 allocs/op
BenchmarkNoEllipsis0-8                	962382211	         1.259 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1-8                	953363317	         1.256 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis10-8               	955995655	         1.260 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis30-8               	957052911	         1.253 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis100-8              	949933554	         1.259 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000-8             	955583925	         1.258 ns/op	       0 B/op	       0 allocs/op
BenchmarkNoEllipsis1000000-8          	945737022	         1.255 ns/op	       0 B/op	       0 allocs/op
PASS
	github.com/cdfmlr/ellipsis	coverage: 89.4% of statements
ok  	github.com/cdfmlr/ellipsis	643.521s
```

`runes = []rune(s[:n * utf8CharMaxSize])`:

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
