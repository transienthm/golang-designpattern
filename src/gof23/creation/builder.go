package creation

type Character struct {
	Name string
	Arms string
}

func (p *Character) SetName(name string) {
	p.Name = name
}

func (p *Character) SetArms(arms string) {
	p.Arms = arms
}

func (p Character) GetArms() string {
	return p.Arms
}

func (p Character) GetName() string {
	return p.Name
}

type Builder interface {
	SetName(name string) Builder
	SetArms(arms string) Builder
	Build() *Character
}

type CharacterBuilder struct {
	character *Character
}

func (p *CharacterBuilder) SetName(name string) Builder {
	if p.character == nil {
		p.character = &Character{}
	}
	p.character.SetName(name)
	return p
}

func (p *CharacterBuilder) SetArms(arms string) Builder {
	if p.character == nil {
		p.character = &Character{}
	}
	p.character.SetArms(arms)
	return p
}

func (p *CharacterBuilder) Build() *Character {
	return p.character
}

type Director struct {
	Builder Builder
}

func (p Director) Create(name string, arms string) *Character {
	return p.Builder.SetName(name).SetArms(arms).Build()
}