# API em Go

## Descrição
Este projeto é uma API RESTful em Go que permite gerenciar produtos. A API interage com um banco de dados PostgreSQL para armazenar as informações dos produtos.

## Pré-requisitos
* Docker
* Go 
* Um cliente HTTP (como Postman) para testar a API

## Instalação
1. **Clone o repositório:**.   
   ```
   git clone [URL]
   ```
2. **Entre no projeto e construa as imagens necessárias do docker**
   ```
   docker-compose up --build -d
   ```
3. **Para verificar se os containers foram iniciados com sucesso, utilize o comando:**
   ```
   docker ps
   ```

## A aplicação estará disponível em http://localhost:8080.
