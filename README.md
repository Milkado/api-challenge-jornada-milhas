# api-challenge-jornada-milhas

## Instalation

`git clone git@github.com:Milkado/api-challenge-jornada-milhas.git`

`cd api-challenge-jornada-milhas`

`go mod download`

`go generate ./ent`

`cp .env_exemplae .env (add DB info)`

`install atlas from: https://atlasgo.io/`

`run one of the two:`

`go run commands/rung.go --command=migrate`
`or`
`atlas migrate apply --dir "file://ent/migrate/migrations" --url "mysql://root:pass@host:port/db"`

`atlas works with MariaDB, PostgreSQL or SQLite`

#### Finally
`go run main.go`

By default the server runs on 8080 port
