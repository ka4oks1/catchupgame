package personell

func init() {

}

type Character interface {
	Move()
}

type Enemy struct {
	Person
}

func (e *Enemy) FindShortWay() {

}

type Person struct {
}

func (p *Person) Move() {

}
