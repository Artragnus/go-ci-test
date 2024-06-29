# GO CI Test

Este é um projeto básico em Go que demonstra um fluxo de integração contínua usando Docker e GitHub Actions para executar testes e analisar a cobertura de código com o SonarCloud.

## Estrutura do Projeto

```plaintext
.
├── Dockerfile
├── README.md
├── coverage.out
├── math.go
├── math_test.go
├── sonar-project.properties
└── .github
    └── workflows
        └── ci.yml
```

## Configuração do Projeto

### Dockerfile
O Dockerfile é usado para criar uma imagem Docker que configura o ambiente para executar o código Go e os testes.

### math.go
Este arquivo contém a implementação das funções matemáticas que estão sendo testadas.

### math_test.go
Este arquivo contém os testes para as funções definidas em `math.go`.

### sonar-project.properties
Este arquivo contém as configurações para a integração com o SonarCloud.

### .github/workflows/ci.yml
Este arquivo define o fluxo de trabalho do GitHub Actions para CI. Ele executa os testes e envia os resultados de cobertura para o SonarCloud.

## Executando o Projeto

### Pré-requisitos
- Docker
- Go
- Conta no SonarCloud

### Passos para Executar

1. **Construir a imagem Docker:**

```bash
docker build -t meu-projeto-go .
```

2. **Executar os testes:**

```bash
docker run --rm meu-projeto-go go test
```

3. **Gerar a cobertura de testes:**

```bash
docker run --rm meu-projeto-go go test -coverprofile=coverage.out
```

4. **Análise automática pelo SonarCloud através do GitHub Actions:**


### Workflow do GitHub actions


```yaml
name: ci-golang-workflow
on:
  pull_request:
    branches:
      - develop
jobs:
  check-application:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: "1.15"
      - name: Run go test and coverage
        run: go test -coverprofile=coverage.out
      - name: SonarCloud Scan
        uses: sonarsource/sonarcloud-github-action@v2.3.0
        env:
          SONAR_TOKEN: ${{ secrets.SONAR_TOKEN }}

  ```



