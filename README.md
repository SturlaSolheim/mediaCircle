# Media Circle (arbeidstittel)

Dette er et rest api for Media Circle. En applikasjon for roterende medialister

## Hvordan bruke

### Kjør applikasjonen:

go run main.go

### Kjøre tester

For å kjøre alle testene kan du bruke tst.sh scriptet. Det genererer mocks og kjører tester

## Flyt for å legge til endepunkt i APIet

- Legg til det nye endepunktet med tilhørende dto objekter i [openapi.yaml](./api/openapi.yaml) fila. Husk å spesifisere en tag
- Generer objekter og interfacer ved å kjøre generate.sh
- Hvis du har lagt til en tag som ikke har vært brukt før vil det bli generert et nytt interface i generated mappen. Lag en ny handler i handlers mappen som implementerer dette interface
- Denne nye handleren må også legges inn i composite_handler.go og routes.go.
- Så er det bare å begynne å implementere

## TODO

[x] Fikse dependancy management
[x] Enviroment variabler
[x] Forskjellige profiler
[x] Sette opp enhetstester
[x] OpenApi for kontrakt
[x] Implementere agnostisk databasemigrering (brukte gorm automigrate)
