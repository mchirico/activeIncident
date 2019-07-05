


[![Build Status](https://travis-ci.org/mchirico/activeIncident.svg?branch=master)](https://travis-ci.org/mchirico/activeIncident)
[![codecov](https://codecov.io/gh/mchirico/activeIncident/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/activeIncident)
[![Build Status](https://mchirico.visualstudio.com/activeIncident/_apis/build/status/mchirico.activeIncident?branchName=master)](https://mchirico.visualstudio.com/activeIncident/_build/latest?definitionId=7&branchName=master)
[![Board Status](https://mchirico.visualstudio.com/f2f93792-8538-42e3-99ee-11cff3492aa9/876f6067-eebb-428e-a0a4-0fdaa6e6889b/_apis/work/boardbadge/745b312c-827d-41ed-99d6-b0ad733b51b7)](https://mchirico.visualstudio.com/f2f93792-8538-42e3-99ee-11cff3492aa9/_boards/board/t/876f6067-eebb-428e-a0a4-0fdaa6e6889b/Microsoft.RequirementCategory/)

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


