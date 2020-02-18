# easy-db
easy-db is easy way to access database

[![Actions](https://github.com/lyf-coder/easy-db/workflows/CI/badge.svg)](https://github.com/lyf-coder/easy-db/actions?query=workflow%3ACI)
[![GoDoc](https://godoc.org/github.com/lyf-coder/easy-db?status.svg)](https://godoc.org/github.com/lyf-coder/easy-db)
[![Go Report Card](https://goreportcard.com/badge/github.com/lyf-coder/easy-db)](https://goreportcard.com/report/github.com/lyf-coder/easy-db)

## Install

```console
go get github.com/lyf-coder/easy-db/db
```
## Usage
    import (
        "github.com/lyf-coder/easy-db/db"
        "github.com/lyf-coder/easy-db/connect"
    )
    var config = connect.Config{
    	DbType:       "DbType",
    	UserName:     "UserName",
    	Password:     "Password",
    	DatabaseName: "DatabaseName",
    	Host:         "ip",
    	Port:         "27017",
    	Options:      nil,
    }
    
    var db = New(&config)
    
    // exec Db interface func   	