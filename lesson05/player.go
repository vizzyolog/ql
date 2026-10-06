package main

type Player struct {
	HP        int
	Inventory []string
}

func (p *Player) hit(damage int) {
	p.HP -= damage
	if p.HP < 0 {
		p.HP = 0
	}
}

func (p *Player) hasItem(name string) bool {
	for _, item := range p.Inventory {
		if item == name {
			return true
		}
	}
	return false
}

func (p *Player) addItem(name string) {
	if !p.hasItem(name) {
		p.Inventory = append(p.Inventory, name)
	}
}

func (p *Player) removeItem(name string) {
	result := []string{}
	removed := false
	for _, item := range p.Inventory {
		if item == name && !removed {
			removed = true
			continue
		}
		result = append(result, item)
	}
	p.Inventory = result
}
