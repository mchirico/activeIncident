


[![Build Status](https://travis-ci.org/mchirico/activeIncident.svg?branch=master)](https://travis-ci.org/mchirico/activeIncident)
[![codecov](https://codecov.io/gh/mchirico/activeIncident/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/activeIncident)
[![Build Status](https://mchirico.visualstudio.com/activeIncident/_apis/build/status/mchirico.activeIncident?branchName=master)](https://mchirico.visualstudio.com/activeIncident/_build/latest?definitionId=7&branchName=master)

# activeIncident

## Build with vendor
```
export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/script" below
go build -v -mod=vendor ./...
```


## Don't forget golint

```

golint -set_exit_status $(go list ./... | grep -v /vendor/)

```


