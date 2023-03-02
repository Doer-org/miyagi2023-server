package model

import "time"

type User struct {
	ID        string
	Name      string
	Age uint
	Gender Gender
	Birthday time.Time
	Address string
	ProfilIMG string
	Prefecture Prefecture
	CreatedAt time.Time
}

type Gender string
func (g *Gender) String() string {
	return string(*g)
}
const (
	MEN Gender = "MEN"
	WOMEN Gender = "WOMEN"
	OTHER Gender = "OTHER"
)

type Prefecture string
func (p *Prefecture) String() string {
	return string(*p)
}
const (
	MIYAGI Prefecture = "MIYAGI"
	TOYAMA Prefecture = "TOYAMA"
)