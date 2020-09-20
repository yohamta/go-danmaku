package effect

import (
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/yohamta/godanmaku/danmaku/internal/shared"
	"github.com/yohamta/godanmaku/danmaku/internal/sound"

	"github.com/yohamta/godanmaku/danmaku/internal/sprite"
)

type explosion struct{ *baseController }

func (c *explosion) init(e *Effect) {}

func (c *explosion) draw(e *Effect, screen *ebiten.Image) {
	if e.spriteFrame >= sprite.Explosion.Length() {
		return
	}
	sprite.Explosion.SetIndex(e.spriteFrame)
	sprite.Explosion.SetPosition(e.x-shared.OffsetX, e.y-shared.OffsetY)
	sprite.Explosion.Draw(screen)

	// TODO: refactor
	scale := float64(sprite.Explosion.GetWidth()) * e.scale *
		math.Min((1.-(float64(e.spriteFrame)/float64(sprite.Explosion.Length()))+0.5), 1.)
	c.drawGrowEffect(e, scale, scale, 0.5, screen)
}

func (c *explosion) update(e *Effect) {
	if e.updateCount == 0 {
		sound.PlaySe(sound.SeKindBomb)
	}
	if e.updateCount > 0 && e.updateCount%3 == 0 {
		e.spriteFrame++
	}
	if e.spriteFrame >= sprite.Explosion.Length() {
		e.isActive = false
	}
}
