package domain

type Categoria struct {
	nome  string
	Itens []ItemCatalogo
}

func (c *Categoria) Nome() string {
	return c.nome
}

func (c *Categoria) SetNome(nome string) {
	c.nome = nome
}

func (c *Categoria) Icone() string {
	return "📂"
}

func (c *Categoria) Imprimir(indentacao string) {
	println(indentacao + c.Icone() + " " + c.nome)
	for _, item := range c.Itens {
		item.Imprimir(indentacao + "  ")
	}
}
