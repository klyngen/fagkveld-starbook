# Enkelt API + Frontend
Laget for å demonstere hvordan man bruker AzureAD til å autentisere mot API'er

## Oppsett
Set connectionstring til datbase i miljøvariabel
```shell
host=[DATABASE_HOST] port=5432 dbname=postgres user=[USER]@[NAMESPACE] password=[PASSWORD] sslmode=require
```

Kjør backend med
`go run .`

Spinn opp frontend
```shell
yarn; yarn start
```

