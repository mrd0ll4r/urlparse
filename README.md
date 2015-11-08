#urlparse

Benchmarks different methods of url query parsing.

Results:

    BenchmarkParseStdlib-8                    200000              8235 ns/op
    BenchmarkParseStdlibNoUrldecode-8         200000              8525 ns/op
    BenchmarkQueryNew-8                       300000              5460 ns/op
    BenchmarkQueryNewSliced-8                 200000              8015 ns/op
    BenchmarkUrlparseStdlib-8                 200000              8950 ns/op
    BenchmarkPathParseStdlib-8                200000              8105 ns/op
    BenchmarkPathParseStdlibNoUrldecode-8     200000              8130 ns/op
    BenchmarkPathQueryNew-8                   300000              5366 ns/op
    BenchmarkPathQueryNewSliced-8             200000              7955 ns/op
    BenchmarkPathUrlparseStdlib-8             200000              8705 ns/op


The winner is `github.com/chihaya/chihaya/http/query.New()` (`QueryNew` and `PathQueryNew`), but the tradeoff is that every key but `info_hash` is present only once. If the key is found again, the previous value will be overwritten.

As an attempt to solve this, I added `QueryNewSliced`, it's still faster than the stdlib, but way slower than the original.

The stdlibs `net/url.Parse().Values()` (`UrlparseStdlib` and `PathUrlparseStdlib`) is the slowest, because it parses the whole URL, and then the query params.
If we only use `net/url.ParseQuery()` (`ParseStdlib` and `PathParseStdlib`) we are a little bit faster, but not much.

Stripping the URLdecoding from the stdlib implementation turned out to not increase performance (?)