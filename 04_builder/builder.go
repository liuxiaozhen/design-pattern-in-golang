package builder

// Builder
type Builder interface {
	buildPartA()
	buildPartB()
	buildPartC()
}

// Director
type Director struct {
	builder Builder
}

// NewDirector
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct
func (d *Director) Construct() {
	d.builder.buildPartA()
	d.builder.buildPartB()
	d.builder.buildPartC()
}

// builderA
type ConcreteBuilderA struct {
	product string
}

func (a *ConcreteBuilderA) buildPartA() {
	a.product += "sprite|"
}

func (a *ConcreteBuilderA) buildPartB() {
	a.product += "chicken wing|"
}

func (a *ConcreteBuilderA) buildPartC() {
	a.product += "big Mac"
}

func (a *ConcreteBuilderA) GetProduct() string {
	return a.product
}

// builderB
type ConcreteBuilderB struct {
	product string
}

func (b *ConcreteBuilderB) buildPartA() {
	b.product += "cola|"
}

func (b *ConcreteBuilderB) buildPartB() {
	b.product += "chips|"
}

func (b *ConcreteBuilderB) buildPartC() {
	b.product += "McChicken"
}
func (b *ConcreteBuilderB) GetProduct() string {
	return b.product
}
