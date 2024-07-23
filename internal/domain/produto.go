package domain

type Produto struct {
	nome string
}

func (p *Produto) Nome() string {
	return p.nome
}

func (p *Produto) SetNome(name string) {
	p.nome = name
}

func (p *Produto) Icone() string {
	return "📄"
}

func (p *Produto) Imprimir(indentacao string) {
	println(indentacao + p.Icone() + " " + p.nome)
}
