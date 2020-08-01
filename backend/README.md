# backend

```sh
# run
go get
go build
./backend

# install sqlboiler w/postgresql locally
GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

# update db orm from schema
go generate
```

# todo
- health checks
- db
- verify login password
