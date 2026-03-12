# 🔐 AWS VPC Docker Go WebApp

    Aplicação web desenvolvida em Go que gera senhas seguras de 20 caracteres
    utilizando geração criptograficamente segura com o pacote crypto/rand.

    O objetivo do projeto é demonstrar na prática conceitos de:
    - Arquitetura Web + Backend separados
    - Comunicação entre serviços
    - Organização de projetos em Go
    - Containerização com Docker
    - Deploy planejado para AWS EC2
    - Isolamento de serviços com VPC e Security Groups

    Projeto desenvolvido para a disciplina de Serviços em Nuvem.

## 🧠 Arquitetura da Aplicação

    A aplicação é composta por dois serviços independentes.

    Browser  
    ↓  
    Web Server (Go + Gin) — Porta 8080  
    ↓  
    Backend API (Go + Gin) — Porta 25000

    Fluxo da requisição:
    Usuário → Frontend → /api/generate → Backend → Senha gerada

    O frontend atua como proxy interno, evitando problemas de CORS e
    permitindo que o backend permaneça isolado.

## 📂 Estrutura do Projeto

    aws-vpc-docker-go-webapp/             # Raiz do projeto  
    │  
    ├── backend/                          # API responsável pela geração de senhas  
    │   │ 
    │   ├── controllers/                  # Camada HTTP da aplicação  
    │   │   └── password_controller.go    # Endpoint /generate  
    │   │                                  
    │   │  
    │   ├── services/                     # Lógica de negócio  
    │   │   └── password_service.go       # Algoritmo de geração de senha  
    │   │                                 
    │   │  
    │   ├── Dockerfile                    # Container Docker da API  
    │   ├── main.go                       # Inicialização do servidor Gin (porta 25000)  
    │   ├── go.mod                        # Dependências do módulo Go  
    │   └── go.sum  
    │  
    ├── web/                              # Aplicação Web  
    │   │
    │   ├── controllers/                  # Controladores da aplicação web  
    │   │   ├── api_controller.go         # Proxy para o backend (/api/generate)  
    │   │   └── page_controller.go        # Renderização da página HTML  
    │   │  
    │   ├── static/                       # Arquivos estáticos  
    │   │   ├── css/style.css             # Estilização da interface  
    │   │   ├── js/app.js                 # Lógica do frontend (Fetch API)  
    │   │   └── img/icon.png              # Ícone da aplicação  
    │   │  
    │   ├── templates/index.html          # Interface da aplicação  
    │   ├── Dockerfile                    # Container Docker da aplicação web  
    │   ├── main.go                       # Inicialização do servidor web (porta 8080)  
    │   ├── go.mod  
    │   └── go.sum  
    │  
    ├── .gitignore  
    ├── LICENSE  
    └── README.md

## ⚙️ Tecnologias Utilizadas

    - Go 1.25
    - Gin Framework
    - HTML
    - CSS
    - JavaScript
    - Docker
    - AWS EC2
    - AWS VPC
    - Security Groups

## 🔐 Geração de Senhas Seguras

    A geração de senhas utiliza o pacote crypto/rand
    Isso garante geração criptograficamente segura, evitando padrões previsíveis.

    Exemplo de resposta da API:
    {"password": "A$7d@9P!kL2#x8Q%t1Wz"}

## 🌐 Frontend

    O frontend foi desenvolvido utilizando:
    - Go + Gin
    - Templates HTML
    - JavaScript (Fetch API)

    Interface inspirada na estética hacker, utilizando cores verde e preto.

    Funcionalidades:
    - gerar senha segura
    - copiar senha para área de transferência
    - comunicação com backend via proxy

## ☁️ Arquitetura Planejada na AWS

    Internet
    ↓
    EC2 Web Server — Porta 8080 (pública)
    ↓
    EC2 Backend API — Porta 25000 (privada)

    Regras de segurança:
    8080 → acesso público  
    25000 → acesso apenas da EC2 Web

    Isso garante que somente o servidor web possa acessar o backend.

## 🧪 Endpoints

    /generate → Gera senha segura  
    /api/generate → Proxy da aplicação web

## 📎 Autor

    Este projeto foi desenvolvido por Pedro Henrique Leite

## 📄 Licença

    Este projeto é de uso educacional, sem fins comerciais.
    Sinta-se à vontade para utilizar como referência em seus estudos!
