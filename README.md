# iti-backend-challenge

# Solução
Para solucionar o problema a aplicação foi quebrada em 3 partes: 
 
- Infra: Rotas para validação das senhas passadas ( internal/routes/routes.go )
- Negócio: Um arquivo para validações das regras de uma senha forte ( internal/password/password.go )
- Componentes Uteis: Um arquivo para lógicas envolvendo strings ( pkg/strutil/strutil.go ). Que é candidato a ser exposto para outros projetos
 

## Execução
A aplicação pode ser executada de duas maneiras
1. Local
    - Necessária instalação do Golang 1.13
    - Utilizar o determinado comando go run cmd/main/main.go na raiz do projeto
2. Docker
    - A aplicação possui um Dockerfile e sua execução pode ser realizada utilizando os seguintes comandos
        - docker build -f Dockerfile -t iti-challenge .
        -  docker run --net host iti-challenge


## Requisições
A aplicação aceita apenas um POST e o body de um json com a senha a ser validada. Exemplo: 
```
curl -d '{"password":"AbTp9!foo"}' -H "Content-Type: application/json" -X POST http://localhost:3000/users/passwords/validate
```
 
## Respostas 
O sistema possui 3 respostas possíveis. 
1. 200 OK {"IsValid":true,"Errs":null}
2. 200 OK {"IsValid":false,"Errs":<Array com lista de erros da senha>}
3. 400 Bad Request  Malformed Body  [ Quando é enviado um json quebrado , por exemplo  {"password":}]
 
## Testes
Para rodar os testes é necessário utilizar o comando na raiz do projeto:
```   
go test $(go list ./...) 
```

Para rodar testes especificos, é só ir até o diretório desejado e utilizar o comando
```   
go test
```

Os testes unitários da aplicação são: 
- internal/password/validator_test.go
- pkg/strutil/strutil_test.go

Os testes de integração são:
- internal/routes/routes_test.go


## Futuras melhorias
- Criar um fluxo de Chain para interromper e guardar os erros. 
Podendo assim dar um feedback para usuário do que falta na sua senha
- Testes que abrangem as mensagens de erro

