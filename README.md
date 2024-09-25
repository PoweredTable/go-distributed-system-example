# Sobre o projeto
O projeto tem como objetivo apresentar o conceito de programação distribuída através de um simples projeto. 

O projeto é um contador de caracteres de um texto, que é dividido em N partes (`go routines`) e distribuído para N servidores.

# Utilização
Para executar o projeto, execute tanto o lado do `cliente` quanto do `servidor` através do arquivo `main.go`.

## Execução do cliente
Execute o lado cliente utilizando a seguinte linha de comando.
```bash
go run main.go client [file] [routines] [servers...]
```
```bash
go run main.go client example.txt 8 "localhost:8000" "localhost:8001"
```

## Execução do servidor
Execute o lado servidor utilizando a seguinte linha de comando.
```bash
go run main.go server [address] [port]
```
```bash
go run main.go server "localhost" 8000
```
Você pode abrir múltiplos servidores utilizado portas diferentes, basta realizar a execução por diferentes terminais.
