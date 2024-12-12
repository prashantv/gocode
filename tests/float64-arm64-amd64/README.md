## float64: amd64 vs arm64

A repro that shows how float64 operations can return different results
on amd64 vs arm64 ().

Go issue: https://github.com/golang/go/issues/36536

### Relevant output

```
GOARCH=amd64 go test
N - (N / 1e4 * 1e4) =  1.1641532182693481e-10

GOARCH=arm64 go test
N - (N / 1e4 * 1e4) =  6.639311322942376e-11
```
