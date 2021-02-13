sha256_resolver
===============
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/pkgsite.svg)][goRef]

resolve sha256 by brute-force

* speed example @ Thinkpad X1 Extreme 2018 (Core i7-8750H CPU @ 2.20GHz) + WSL2 (Ubuntu20.04) + go1.16beta1
| target | time | 
|-------:|:------|
| a | 0:00.03 |
| ab | 0:00.03 |
| abc | 0:00.41 |
| abcd | 1:49.97 |
| 🤔 | 2:45.27 |

Usage:
------

```
sha256_resolver ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad
```
or
```
sha256-resolver $(echo -n abc | sha256sum | cut -c-64)
```

See [GoDoc][goRef]

Requirements:
-------------
+ go

Install:
--------
+ go1.6 or higher
```
go install github.com/0Delta/sha256_resolver@latest
```

+ go1.5 or lower
```
go get -u github.com/0Delta/sha256_resolver
```

license:
--------
MIT

Author:
-------
0Δ(0deltast@gmail.com)


[goRef](https://pkg.go.dev/github.com/0Delta/sha256_resolver)
