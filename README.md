<h2>Crypto  Upvote</h2>

<p>Este projeto consiste em uma API REST feita com Golang e um frontend feito em React, que disponibiliza um serviço de upvote onde podemos votar em algumas criptomoedas famosas como Bitcoin, Ethereum, Litecoin, Dogecoin, etc.<p>

<p>Temos duas opções para executar o projeto, sendo a primeira via Docker e a segunda seria subindo todos os serviços um a um. De qualquer forma, é importante criar e preencher os arquivos <b>.env</b> dentro da pasta dos serviços. Existem arquivos exemplos dentro dos diretórios (<b>.env.example</b>).</p>

<br/>
<h3>Rodando o app via Docker</h3>

<p>Para rodar o projeto completo via docker-compose, basta executar dois comandos. Os comandos devem ser executados a partir da raiz do projeto, sendo o primeiro para buildar as imagens do front e back:</p>

    docker-compose build

<p>Na sequencia basta executar o comando que vai levantar todos os containers a partir das imagens criadas:</p>

    docker-compose up -d

<p>Com a aplicação rodando, basta entrar no endereço <b>http://localhost:3000/</b> através do navegador. Caso queira testar apenas a API via endpoints, dentro do diretório do backend, existe um arquivo com a collection POSTMAN de endpoints.</p>

<br/>
<h3>Rodando o app separadamente</h3>

<p>Para rodar o app separadamente, vamos precisar de uma instância do banco de dados relacional Postgres rodando na máquina, NodeJs e Golang instalados.</p>

<p>Para rodar o frontend, basta entrar no diretorio do frontend e executar o seguinte comando:</p>

    npm install

<p>esse vai instalar as dependências do projeto, e por fim o comando que irá subir o projeto:

    npm start

<p>Para rodar o backend, basta entrar no diretorio backend e executar o primeiro comando que instala as dependências do projeto:</p>

    go get ./...

<p>Na sequencia o comando que irá buildar o binário da aplicação:</p>

    go build -o ./build

<p>Por fim, o comando que vai subir a API</p>

    ./build

<p>Com a aplicação rodando, basta entrar no endereço <b>http://localhost:3000/</b> através do navegador
