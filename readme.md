# Ponderada 1

## 1. Introdução

Esse repositório contem um código em go que simula Sensor de Radiação Solar Sem Fio HOBOnet RXW-LIB-900. O código utiliza da tecnologia de mqtt para simular o envio dos dados.

## 2. Dependências

para baixar as dependências do projeto, execute o comando abaixo:

```bash
go mod tidy
```

## 3. Execução

O projeto aproveita das goroutines para que o broker, publisher e subscriber sejam rodados simultaneamente. Para executar o projeto, execute o comando abaixo:

```bash
go run *.go
```

## 4. Vídeo

O vídeo de apresentação do projeto pode ser encontrado [aqui](https://youtu.be/c8rqMAsyQTQ).