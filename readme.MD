# Load Tester

Um sistema CLI em Go para realizar testes de carga em um serviço web. Permite ao usuário fornecer a URL do serviço, o número total de requests e a quantidade de chamadas simultâneas. O sistema gera um relatório detalhado após a execução dos testes.

## Funcionalidades

- Realiza requests HTTP para a URL especificada.
- Distribui os requests de acordo com o nível de concorrência definido.
- Garante que o número total de requests seja cumprido.
- Gera um relatório contendo:
    - Tempo total gasto na execução.
    - Quantidade total de requests realizados.
    - Quantidade de requests com status HTTP 200.
    - Distribuição de outros códigos de status HTTP (como 404, 500, etc.).

## Como Usar

### Pré-requisitos

- Docker instalado em sua máquina.

### Construir a Imagem Docker

Para construir a imagem Docker, execute o seguinte comando:

```bash
make build
```

Para rodar: 
```bash
make run