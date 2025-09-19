# Como Rodar o Projeto

Para iniciar o projeto, você precisa ter o Docker, Docker Compose e o sqlx instalados. O projeto foi configurado como um monorepo, e todas as dependências de infraestrutura são gerenciadas pelo docker-compose.yml na raiz.

Caso ainda não possua, baixe e instale o sqlx-cli para gerenciar as migrações do banco de dados.

1. Configurar o Ambiente

Primeiro, crie seu arquivo de variáveis de ambiente. A partir da raiz do projeto, copie o arquivo de exemplo e, em seguida, preencha-o com as suas credenciais.

```Bash

cp .env.example .env

```

2. Inicializar a Infraestrutura e o Banco de Dados

Agora, inicie os contêineres do PostgreSQL e do RabbitMQ. Eles serão executados em segundo plano.


```Bash

docker-compose up -d

```

Em seguida, exporte a variável de ambiente DATABASE_URL no seu terminal e rode as migrações do banco de dados. Isso irá criar as tabelas necessárias para a aplicação.

```Bash

export DATABASE_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"
sqlx migrate run

```


3. Executar a Aplicação Go

Com a infraestrutura e o banco de dados configurados, você pode iniciar o serviço da API Go.

Vá para o diretório go-api e execute o comando:

```Bash

go run cmd/main.go

```

Isso irá iniciar o servidor web. Se tudo estiver correto, você verá a mensagem de sucesso no seu terminal, e sua API estará pronta para receber requisições em http://localhost:8080.

Próximos Passos

Se você precisar parar todos os contêineres e remover os volumes associados, vá para a raiz do projeto e execute:

```Bash

docker-compose down --volumes

```