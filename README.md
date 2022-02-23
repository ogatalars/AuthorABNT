# AuthorABNT

Projeto que visa transformar o nome de uma pessoa em formato ABNT.

O projeto precisa de um arquivo CSV com os nomes a serem normalizados em formato ABNT de forma sequêncial e retorna um CSV na pasta ```AuthorsABNT``` o arquivo ```Autores_ABNT.csv``` com o sequênciamento:

-> ```Nome Autor``` ; ```ABNT Longo``` ; ```ABNT sem ponto``` e ```ABNT com ponto```

## Run
Nada além de garantir que a pasta ```NameCSV``` esteja alimentada com o arquivo ```Authors``` com uma única coluna com a sequência de nomes que devem ser normalizados.

```go run main.go```
