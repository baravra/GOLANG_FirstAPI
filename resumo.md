# AULA DE GO E DOCKER

## LISTA DE COMANDOS PARA GO

Inicializar o projeto:
```go mod init {nomeProjeto}```

Instalar pacotes
```go get {link}```

Rodar server:
```go run {arquivo}```

Rodar o container:
```docker compose up -d {nomeContainer}```
OBS: é preciso que esteja rodando no mesmo caminho de onde esta o arquivo docker-compose.yml e que o docker desktop esteja aberto

Gerar imagem da aplicação
```docker build -t {nomeDaImagem} .```

## OBSERVAÇÕES
Pra uma struct ser visivel para os outros pacotes a primeira letra deve ser maiuscula
