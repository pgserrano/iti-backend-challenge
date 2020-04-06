# iti-backend-challenge

# Solução
Para solucionar o problema a aplicação foi quebrada em 3 partes: 
 
- Infra: Rotas para validação das senhas passadas ( internal/routes/routes.go )
- Negócio: Um arquivo para validações das regras de uma senha forte ( internal/password/password.go )
- Componentes Uteis: Um arquivo para lógicas envolvendo strings ( pkg/strutil/strutil.go ). Que é candidato a ser exposto para outros projetos
 

# Utilização
##Execução
A aplicação pode ser executada de duas maneiras
1. Local
    - Necessária instalação do Golang 1.13
    - Utilizar o determinado comando go run cmd/main/main.go na raiz do projeto
2. Docker
    - A aplicação possui um Dockerfile e sua execução pode ser realizada utilizando os seguintes comandos
        - docker build -f Dockerfile -t iti-challenge .
        -  docker run --net host iti-challenge
##Requisições
A aplicação aceita apenas um POST e o body com a senha a ser validada. Exemplo: 
```
curl -d 'AbTp9!foo' -X POST http://localhost:3000/users/passwords/validations/isValid
```
 
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
- internal/password/password_test.go
- pkg/strutil/strutil_test.go

Os testes de integração são:
- internal/routes/routes_test.go


## Futuras melhorias
- Criar um fluxo de Chain para interromper e guardar os erros. 
Podendo assim dar um feedback para usuário do que falta na sua senha


