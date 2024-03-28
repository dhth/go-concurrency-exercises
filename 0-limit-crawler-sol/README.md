# Limit your crawler

## Logs

Original

```
[48:10] fetching for url: http://golang.org/, depth: 4
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] found: http://golang.org/ "The Go Programming Language"
[48:10] fetching for url: http://golang.org/cmd/, depth: 3
[48:10] url: http://golang.org/cmd/ - fetcher returned 0 urls
not found: http://golang.org/cmd/
[48:10] fetching for url: http://golang.org/pkg/, depth: 3
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] found: http://golang.org/pkg/ "Packages"
[48:10] fetching for url: http://golang.org/pkg/os/, depth: 2
[48:10] url: http://golang.org/pkg/os/ - fetcher returned 2 urls
[48:10] found: http://golang.org/pkg/os/ "Package os"
[48:10] fetching for url: http://golang.org/cmd/, depth: 2
[48:10] url: http://golang.org/cmd/ - fetcher returned 0 urls
not found: http://golang.org/cmd/
[48:10] fetching for url: http://golang.org/pkg/fmt/, depth: 2
[48:10] url: http://golang.org/pkg/fmt/ - fetcher returned 2 urls
[48:10] fetching for url: http://golang.org/, depth: 1
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] found: http://golang.org/ "The Go Programming Language"
[48:10] fetching for url: http://golang.org/pkg/, depth: 1
[48:10] fetching for url: http://golang.org/pkg/, depth: 0
[48:10] found: http://golang.org/pkg/fmt/ "Package fmt"
[48:10] fetching for url: http://golang.org/pkg/, depth: 1
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] found: http://golang.org/pkg/ "Packages"
[48:10] fetching for url: http://golang.org/pkg/os/, depth: 0
[48:10] fetching for url: http://golang.org/, depth: 2
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] fetching for url: http://golang.org/pkg/fmt/, depth: 0
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] found: http://golang.org/ "The Go Programming Language"
[48:10] fetching for url: http://golang.org/cmd/, depth: 0
[48:10] fetching for url: http://golang.org/cmd/, depth: 1
[48:10] fetching for url: http://golang.org/cmd/, depth: 0
[48:10] fetching for url: http://golang.org/, depth: 0
[48:10] fetching for url: http://golang.org/, depth: 1
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] found: http://golang.org/ "The Go Programming Language"
[48:10] url: http://golang.org/cmd/ - fetcher returned 0 urls
not found: http://golang.org/cmd/
[48:10] fetching for url: http://golang.org/pkg/, depth: 1
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] found: http://golang.org/pkg/ "Packages"
[48:10] fetching for url: http://golang.org/cmd/, depth: 0
[48:10] fetching for url: http://golang.org/pkg/, depth: 0
[48:10] fetching for url: http://golang.org/pkg/os/, depth: 0
[48:10] found: http://golang.org/pkg/ "Packages"
[48:10] fetching for url: http://golang.org/pkg/os/, depth: 0
[48:10] fetching for url: http://golang.org/pkg/fmt/, depth: 0
[48:10] fetching for url: http://golang.org/, depth: 0
[48:10] fetching for url: http://golang.org/, depth: 0
[48:10] fetching for url: http://golang.org/cmd/, depth: 0
[48:10] fetching for url: http://golang.org/pkg/fmt/, depth: 0
[48:10] fetching for url: http://golang.org/cmd/, depth: 0
```

Rate Limited

```
[49:48] ping!
[49:48] fetching for url: http://golang.org/, depth: 4
[49:48] url: http://golang.org/ - fetcher returned 2 urls
[49:48] found: http://golang.org/ "The Go Programming Language"
[49:48] fetching for url: http://golang.org/cmd/, depth: 3
[49:48] fetching for url: http://golang.org/pkg/, depth: 3
[49:49] ping!
[49:49] url: http://golang.org/cmd/ - fetcher returned 0 urls
not found: http://golang.org/cmd/
[49:50] ping!
[49:50] url: http://golang.org/pkg/ - fetcher returned 4 urls
[49:50] found: http://golang.org/pkg/ "Packages"
[49:50] fetching for url: http://golang.org/pkg/os/, depth: 2
[49:50] fetching for url: http://golang.org/, depth: 2
[49:50] fetching for url: http://golang.org/cmd/, depth: 2
[49:50] fetching for url: http://golang.org/pkg/fmt/, depth: 2
[49:51] ping!
[49:51] url: http://golang.org/pkg/os/ - fetcher returned 2 urls
[49:51] found: http://golang.org/pkg/os/ "Package os"
[49:51] fetching for url: http://golang.org/pkg/, depth: 1
[49:51] fetching for url: http://golang.org/, depth: 1
[49:52] ping!
[49:52] url: http://golang.org/ - fetcher returned 2 urls
[49:52] found: http://golang.org/ "The Go Programming Language"
[49:52] fetching for url: http://golang.org/cmd/, depth: 1
[49:52] fetching for url: http://golang.org/pkg/, depth: 1
[49:53] ping!
[49:53] url: http://golang.org/cmd/ - fetcher returned 0 urls
not found: http://golang.org/cmd/
[49:54] ping!
[49:54] url: http://golang.org/pkg/fmt/ - fetcher returned 2 urls
[49:54] found: http://golang.org/pkg/fmt/ "Package fmt"
[49:54] fetching for url: http://golang.org/, depth: 1
[49:54] fetching for url: http://golang.org/pkg/, depth: 1
[49:55] ping!
[49:55] url: http://golang.org/pkg/ - fetcher returned 4 urls
[49:55] found: http://golang.org/pkg/ "Packages"
[49:55] fetching for url: http://golang.org/pkg/os/, depth: 0
[49:55] fetching for url: http://golang.org/pkg/fmt/, depth: 0
[49:55] fetching for url: http://golang.org/cmd/, depth: 0
[49:55] fetching for url: http://golang.org/, depth: 0
[49:56] ping!
[49:56] url: http://golang.org/ - fetcher returned 2 urls
[49:56] found: http://golang.org/ "The Go Programming Language"
[49:56] fetching for url: http://golang.org/pkg/, depth: 0
[49:56] fetching for url: http://golang.org/cmd/, depth: 0
[49:57] ping!
[49:57] url: http://golang.org/cmd/ - fetcher returned 0 urls
not found: http://golang.org/cmd/
[49:58] ping!
[49:58] url: http://golang.org/pkg/ - fetcher returned 4 urls
[49:58] found: http://golang.org/pkg/ "Packages"
[49:58] fetching for url: http://golang.org/pkg/os/, depth: 0
[49:58] fetching for url: http://golang.org/, depth: 0
[49:58] fetching for url: http://golang.org/pkg/fmt/, depth: 0
[49:58] fetching for url: http://golang.org/cmd/, depth: 0
[49:59] ping!
[49:59] url: http://golang.org/ - fetcher returned 2 urls
[49:59] found: http://golang.org/ "The Go Programming Language"
[49:59] fetching for url: http://golang.org/cmd/, depth: 0
[49:59] fetching for url: http://golang.org/pkg/, depth: 0
[50:00] ping!
[50:00] url: http://golang.org/pkg/ - fetcher returned 4 urls
[50:00] found: http://golang.org/pkg/ "Packages"
[50:00] fetching for url: http://golang.org/pkg/os/, depth: 0
[50:00] fetching for url: http://golang.org/, depth: 0
[50:00] fetching for url: http://golang.org/pkg/fmt/, depth: 0
[50:00] fetching for url: http://golang.org/cmd/, depth: 0
[50:01] done; quitting timer
```

## Just the fetch logs

Original

```
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] url: http://golang.org/cmd/ - fetcher returned 0 urls
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] url: http://golang.org/pkg/os/ - fetcher returned 2 urls
[48:10] url: http://golang.org/cmd/ - fetcher returned 0 urls
[48:10] url: http://golang.org/pkg/fmt/ - fetcher returned 2 urls
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] url: http://golang.org/ - fetcher returned 2 urls
[48:10] url: http://golang.org/cmd/ - fetcher returned 0 urls
[48:10] url: http://golang.org/pkg/ - fetcher returned 4 urls
```

Rate Limited

```
[49:48] url: http://golang.org/ - fetcher returned 2 urls
[49:49] url: http://golang.org/cmd/ - fetcher returned 0 urls
[49:50] url: http://golang.org/pkg/ - fetcher returned 4 urls
[49:51] url: http://golang.org/pkg/os/ - fetcher returned 2 urls
[49:52] url: http://golang.org/ - fetcher returned 2 urls
[49:53] url: http://golang.org/cmd/ - fetcher returned 0 urls
[49:54] url: http://golang.org/pkg/fmt/ - fetcher returned 2 urls
[49:55] url: http://golang.org/pkg/ - fetcher returned 4 urls
[49:56] url: http://golang.org/ - fetcher returned 2 urls
[49:57] url: http://golang.org/cmd/ - fetcher returned 0 urls
[49:58] url: http://golang.org/pkg/ - fetcher returned 4 urls
[49:59] url: http://golang.org/ - fetcher returned 2 urls
[50:00] url: http://golang.org/pkg/ - fetcher returned 4 urls
```
