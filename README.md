# Media Circle (arbeidstittel)
Dette er et rest api for Media Circle. En applikasjon for roterende medialister

## Hvordan bruke
### Kjør applikasjonen:
go run main.go

### Kjøre tester
For å kjøre alle testene kan du bruke test.sh scriptet. Det genererer mocks og kjører tester

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
[] implementere daglig oppdatering med cron
[] sette opp frontend
[] utvide databasemodellen

## Forslag til tags i api
- anmeldelse
- medie
- bruker
- gruppe
- liste
- listeinstans
- notifikasjon


## Møte 11. desember 2025
gjenbruke anmeldselser i nye listeinstanser
bruker kan ikke abbonerer på lister, det må være gjennom en gruppe. Man kan være i gruppe alene
forskjellige typer random?
admin/roller på gruppe
notofikasjoener

kommentere anmeldelser internt i gruppe
reaksjoner på anmeldelser intern i gruppa
fortsatt mulig å se anmeldelser globalt.
fortsatt mulig å se anmeldelser globalt.
private eller public anmeldelser på review nivå
rapportere upassende innhold

import av grupper fra 1001
