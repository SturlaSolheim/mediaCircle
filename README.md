# Media Circle (arbeidstittel)

Dette er et rest api for Media Circle. En applikasjon for roterende medialister

## Hvordan bruke

### Kjør applikasjonen:
go run main.go

### Generere mocks
Når du har endret eller lagt til et nytt interface, kjør denne:
go run github.com/vektra/mockery/v2@latest                                                                                                               │  -d '{"name": "Summer Vibes"}'

### Kjør alle testene:
go test ./...


## TODO

[x] Fikse dependancy management
[x] Enviroment variabler
[x] Forskjellige profiler
[x] Sette opp enhetstester
[] OpenApi for kontrakt
[] Implementere agnostisk databasemigrering
