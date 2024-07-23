package main

import (
	"github.com/br4tech/go-composite-example/internal/domain"
	"github.com/br4tech/go-composite-example/internal/service"
)

func main() {
	camiseta := &domain.Produto{}
	camiseta.SetNome("Camiseta")

	calca := &domain.Produto{}
	calca.SetNome("Calça")

	tenis := &domain.Produto{}
	tenis.SetNome("Tênis")

	// Criação das categorias usando o método SetNome
	roupas := &domain.Categoria{}
	roupas.SetNome("Roupas")
	roupas.Itens = []domain.ItemCatalogo{camiseta, calca}

	esporte := &domain.Categoria{}
	esporte.SetNome("Esporte")
	esporte.Itens = []domain.ItemCatalogo{tenis}

	// Criação do catálogo principal usando o método SetNome
	catalogo := &domain.Categoria{}
	catalogo.SetNome("Catálogo")
	catalogo.Itens = []domain.ItemCatalogo{roupas, esporte}

	service := service.NewCatalogoService(catalogo)
	service.ImprimirCatalogo()
}
