package service

import (
	"fmt"

	"github.com/br4tech/go-composite-example/internal/domain"
)

type CatalogoService struct {
	catalogo domain.ItemCatalogo
}

func NewCatalogoService(catalogo domain.ItemCatalogo) *CatalogoService {
	return &CatalogoService{catalogo: catalogo}
}

func (cs *CatalogoService) AdicionarItem(item domain.ItemCatalogo) {
	if categoria, ok := cs.catalogo.(*domain.Categoria); ok {
		categoria.Imprimir("")
	} else {
		fmt.Println("Erro: Não é possível adicionar itens à raiz do catálogo.")
	}
}

func (cs *CatalogoService) ImprimirCatalogo() {
	cs.catalogo.Imprimir("")
}
