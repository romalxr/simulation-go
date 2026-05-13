package entity

type Creature interface {
	Occupier
	Speed() int
	Cooldown() int
	SetCooldown(int)
	Hp() int
	SetHp(int)
}
