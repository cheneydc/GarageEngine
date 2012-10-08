package SpaceCookies

import (
	. "github.com/vova616/GarageEngine/Engine"
	"math/rand"
	"time"
)

type Power int

type PowerUp struct {
	BaseComponent
	Type Power
}

const (
	Speed  = Power(2)
	Damage = Power(6)
	Range  = Power(1)
	HP     = Power(5)
)

func NewPowerUp(typ Power) *PowerUp {
	return &PowerUp{BaseComponent: NewComponent(), Type: typ}
}

func CreatePowerUp(position Vector) {
	chance := rand.Int() % 100
	if chance <= 2 {
		c := PowerUpGO.Clone()
		c.Transform().SetParent2(GameSceneGeneral.Layer2)
		c.Transform().SetPosition(position)

		index := (rand.Int() % 6)

		for index == 2 || index == 3 || index == 4 {
			index = (rand.Int() % 6)
		}

		index += 6

		c.Sprite.SetAnimationIndex(int(index))

		c.AddComponent(NewPowerUp(Power(index - 5)))
	} else if chance <= 5 {
		c := PowerUpGO.Clone()
		c.Transform().SetParent2(GameSceneGeneral.Layer2)
		c.Transform().SetPosition(position)

		index := int(HP) - 1

		for index == 2 || index == 3 {
			index = (rand.Int() % 6)
		}

		index += 6

		c.Sprite.SetAnimationIndex(int(index))

		c.AddComponent(NewPowerUp(Power(index - 5)))
	}

}

func (pu *PowerUp) OnCollisionEnter(arbiter *Arbiter) bool {
	if pu.GameObject() != nil && (arbiter.GameObjectA() == Player || arbiter.GameObjectB() == Player) {
		switch pu.Type {
		case Speed:
			PlayerShip.Speed += 30000
		case Damage:
			var dmg *DamageDealer
			dmg = PlayerShip.Missle.GameObject().ComponentTypeOfi(dmg).(*DamageDealer)
			dmg.Damage += 50
		case Range:
			var dst *Destoyable
			dst = PlayerShip.Missle.GameObject().ComponentTypeOfi(dst).(*Destoyable)
			dst.aliveDuration += time.Millisecond * 500
		case HP:
			PlayerShip.Destoyable.HP = PlayerShip.Destoyable.FullHP
			PlayerShip.OnHit(nil, nil)
		}
		pu.GameObject().Destroy()
	}
	return true
}