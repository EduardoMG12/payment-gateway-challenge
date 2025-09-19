# Como Rodar o Projeto

Para iniciar o projeto, a forma mais rápida e recomendada é usar o script de setup. Ele automatiza a configuração do ambiente, a instalação das dependências e a preparação do banco de dados para desenvolvimento.

Se você não tem Docker e Docker Compose instalados, instale-os antes de começar. Se você também não tiver o sqlx-cli instalado, ele será necessário para as migrações do banco de dados.

1. Setup Rápido para Desenvolvedores

Primeiro, torne o script executável. A partir da raiz do projeto, execute o comando:
```Bash

chmod +x scripts/setup.sh

```

Em seguida, rode o script. Ele irá:

    Copiar o arquivo .env.example para .env.

    Instalar as dependências do Node.js.

    Configurar o Husky para os hooks do Git.

```Bash

./scripts/setup.sh

```

Depois de executar o script, abra o arquivo .env para configurar as credenciais do seu banco de dados.

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