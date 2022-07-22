# 讓值固定精度為 2 的做法

條件：

* 沒有計算
* 本來小數點就不會超過兩位

```shell
goos: windows
goarch: amd64
pkg: github.com/marco79423/golang-lab/fixed-precision
cpu: Intel(R) Core(TM) i5-8400 CPU @ 2.80GHz
BenchmarkStringFixedBank
BenchmarkStringFixedBank-6       1761231              7541 ns/op
BenchmarkStringFixed
BenchmarkStringFixed-6           4598810              4129 ns/op
BenchmarkW1
BenchmarkW1-6                    2006449              6395 ns/op
BenchmarkW2
BenchmarkW2-6                    2135181              4984 ns/op
PASS
```
