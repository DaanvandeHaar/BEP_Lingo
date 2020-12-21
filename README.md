
# BEP_Lingo
![Go-Workflow](https://github.com/typical-go/typical-rest-server/workflows/Go/badge.svg)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-79%25-brightgreen.svg?longCache=true&style=flat)</a>
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=DaanvandeHaar_BEP_Lingo&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=DaanvandeHaar_BEP_Lingo)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=DaanvandeHaar_BEP_Lingo&metric=security_rating)](https://sonarcloud.io/dashboard?id=DaanvandeHaar_BEP_Lingo)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=DaanvandeHaar_BEP_Lingo&metric=sqale_index)](https://sonarcloud.io/dashboard?id=DaanvandeHaar_BEP_Lingo)

</br>

Een lingo API voor het vak BEP aan de HU

## API endpoints

| Endpoint            | Prefix | Functie                                                                   |
|---------------------|--------|---------------------------------------------------------------------------|
| /game/new           | /api   | Start een nieuw spel                                                      |
| /game/current/guess | /api   | Speel door op je meest recente spel                                       |
| /game/score         | /api   | Verkrijg een lijst met de hoogste scores                                  |
| /auth/login         |        | Login en verkrijg een JWT token om gebruik te maken van de /api endpoints |
| auth/signup         |        | Maak een nieuw account om de API te kunnen gebruiken                      |
| /jwt                | /test  | Test functie om tijdelijke jwt tokens te genereren                        |
| /getrandom          | /test  | Test functie om een random woord met de gewenste lengte te verkrijgen     |

## Build tools en pipeline
Er is gebruik gemaakt van github actions voor CI/CD. Bij elke merge of push op master op master worden automatische alle tests gedraaid en word er gekeken naar code smells en security risks met SonarCloud. Als deze slagen zal de merge/push succesvol zijn.
Bij elke release word net als bij elke merge of push op master alle unit test gedraaid. Vervolgens word het project getest en vervolgens via Google App Engine (Google Cloud) gebuild en automatisch gedeployeerd.


## Mate van functionaliteit
De lingo app bevat alle gewenste functionaliteit. Er kan een spel gestart en gespeeld worden. Ook kunnen de highscores bekeken worden.

## Testorganistatie
Er is gebruik gemaakt van Unit tests om alle services te testen. Integration tests worden gedraaid via postman. 

✅ TODO: CI met Semaphore

## Clean tests
Er is gebruik gemaakt van clean tests. Test draaien onafhankelijk en testen de daadwerkelijke functies. Bij de tests word een mock repository meegegeven doormiddel van dependency injection. Hierdoor word de production database tijdens testen niet gebruikt. 

## Coverage en mutation testing
Coverage word bijgehouden via go coverage. Bij elke push of merge met master word een nieuwe coverage badge gecreëerd via gopherbadger.

## Mate van structuur
Er is gebruik gemaakt van hexagonal domain driven design, de domein laag en daarbij behorende services zijn opgedeeld in packages met packages die dezelfde structs gebruiken. Ook is er veelvuldig gebruik gemaakt van dependency injection, Repositories worden meegegeven aan services die vervolgens worden meegegeven aan API endpoints. Op deze manier kunnen repositories en services makkelijk verwisseld worden voor andere repositories en services. Dit wordt ook tijdens het testen toegepast waar een mock repository aan de service tests wordt meegegeven.

✅ TODO: Auth page veranderen naar service.

## Mate van netheid

| SOLID Principe           | Toepassing                                                                                                |
|--------------------------|-----------------------------------------------------------------------------------------------------------|
| Single repsponsibillity  | Zoveel mogelijk gebruik gemaakt van kleine functies die maar een taak hebben                              |
| Open/closed              | Gebruik gemaakt van interfaces per service makkelijk uit te breiden zonder de core-functies aan te passen |
| Liskov                   | Geen gebruik gemaakt van overerving                                                                       |
| Interfaces Segregation   | Gebruik gemaakt van interfaces op service en repository niveau                                            |
| Dependency inversion     | Zwaar gebruik van dependency injection in alle services                                                   |

## Static analysis tools
Er is gebruik gemaakt van sonarCloud via CI. Bij elke merge of push op master zal netheid van de code gecontroleerd worden, vervolgens zullen hiervoor “grades” gecreëerd worden. De badges hiervan zullen in README.md geplaatst worden.

## Performance en security analysis
Mogelijke security risks worden gecontroleerd door SonarCloud. Performance is bij te houden via Google cloud App Eninge.

## Deployment
Gedeployed op Google cloud App Engine doormiddel van CD met github actions.

## Creatieve ruimte
Ik heb tijdens dit project bijna alleen maar gebruikt van programmeertalen, technologieën en frameworks die ik voor dit project nog nooit eerder had gebruikt. De learning-curve was vooral in het begin lastig maar ik heb erg veel geleerd. 




