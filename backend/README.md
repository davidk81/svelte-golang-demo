# backend

```sh
# run
go get
go build
./backend

# install sqlboiler w/postgresql locally
GO111MODULE=off go get -u -t github.com/volatiletech/sqlboiler
GO111MODULE=off go get github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

# if that doesnt work, try
go get -u -t -v github.com/volatiletech/sqlboiler
go get -u -v github.com/volatiletech/sqlboiler/drivers/sqlboiler-psql

# install tools for sql boiler test
brew install libpq

# update db orm from schema
(cd ./patientdb && go generate)

# build app as docker image
docker build . -t patient-backend

# run app as docker image
docker run --rm -it -p 8000:8000 patient-backend
```

# todo
- health checks
- db
- verify login password
