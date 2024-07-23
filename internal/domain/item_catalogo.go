package domain

type ItemCatalogo interface {
	Nome() string
	Icone() string // Novo método para retornar o ícone
	Imprimir(indentacao string)
}
