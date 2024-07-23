# 🎨 Design Composite em Go: Construindo Catálogos Flexíveis com 🌳 Árvore Genealógica 👪

Este projeto demonstra a implementação do padrão de projeto Composite em Go, utilizando a analogia de uma árvore genealógica para criar um catálogo de produtos flexível e escalável.

## 🎯 Objetivo

O objetivo deste projeto é ilustrar como o padrão Composite permite:

- **Representar estruturas hierárquicas**: Organizar produtos em categorias e subcategorias de forma intuitiva.
- **Tratar objetos de forma uniforme**: Manipular produtos e categorias com a mesma interface, simplificando o código.
- **Flexibilidade e escalabilidade**: Adicionar novos tipos de itens ao catálogo sem modificar o código existente.

## 🚀 Tecnologias Utilizadas

- **Go**: Linguagem de programação principal.
- **Arquitetura Hexagonal**: Organização do projeto em camadas para melhor modularidade e testabilidade.

## 📂 Estrutura do Projeto

```bash
```

## 💡 Conceitos-chave

- **Composite**: Padrão de projeto estrutural que permite compor objetos em estruturas de árvore e tratar objetos individuais e composições de forma uniforme.
- **Polimorfismo**: Capacidade de um objeto assumir diferentes formas, permitindo que objetos de diferentes tipos sejam tratados como se fossem do mesmo tipo.
- **Interface ItemCatalogo**: Define o contrato que todos os componentes do catálogo (produtos e categorias) devem seguir.

## 🛠️ Como usar

1. Clone o repositório: 
    ```sh
      git clone https://github.com/br4tech/go-composite-example
    ```
2. Instale as dependências:
    ```sh
      go mod download
    ```
3. Execute o programa: 
```bash
  go run cmd/main.go
```

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou pull requests.