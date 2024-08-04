# api-challenge-jornada-milhas

## Instalation

`git clone git@github.com:Milkado/api-challenge-jornada-milhas.git`

`cd api-challenge-jornada-milhas`

`go mod download`

`go generate ./ent`

`cp .env_exemplae .env (add DB info)`

`install atlas from: https://atlasgo.io/`

`run one of the two:`

`go run commands/run.go --command=migrate`
`or`
`atlas migrate apply --dir "file://ent/migrate/migrations" --url "mysql://root:pass@host:port/db"`

`atlas works with MariaDB, PostgreSQL or SQLite`

`generate the jwt secret with:`

`go run commands/run.go --command=generate_secret`

#### Finally
`go run main.go`

By default the server runs on 8080 port

#### Tests

Tests are in the `test` folder, but go runs tests directly from the go path. So, the tests need to be built before running them.

`go test ./test/*_test.go` or `go test ./test`

The first one will build the test by the chosen test, the second one will build all the tests.

