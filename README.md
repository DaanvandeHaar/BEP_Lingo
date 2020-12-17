
# BEP_Lingo
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-79%25-brightgreen.svg?longCache=true&style=flat)</a>
Een lingo API voor het vak BEP aan de HU

## Build tools en pipeline
Er is gebruik gemaakt van github actions voor CD. Bij elke merge met master of push op master worden automatische alle tests gedraaid en word er gekeken naar code smells en security risks met SonarCloud. Als deze slagen zal de merge/push succesvol zijn.

TODO: CD met google cloud

## Mate van functionaliteit
De lingo app bevat alle gewenste functionaliteit. Er kan een spel gestart en gespeeld worden. Ook kunnen de highscores bekeken worden.

## Testorganistatie
Er is gebruik gemaakt van Unit tests om alle services te testen.

TODO: Intergration testing

## Clean tests
Er is gebruik gemaakt van clean tests. Test draaien onafhangkelijk en testen de daadwerkelijke functies. Bij de tests word een mockRepository meegegeven doormiddel van dependency injection. Hierdoor word de production database tijdens testen niet gebruikt. 

