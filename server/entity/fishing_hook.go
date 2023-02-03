package entity

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/cube/trace"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

func NewFishingHook(pos mgl64.Vec3, owner world.Entity) *Ent {
	return Config{Behaviour: fishingHookConf.New(owner)}.New(FishingHookType{}, pos)
}

var fishingHookConf = ProjectileBehaviourConfig{
	Gravity: 0.1,
	Drag:    0.02,
	Hit:     Tick,
}

type FishingHookType struct{}

func (FishingHookType) EncodeEntity() string {
	return "minecraft:fishing_hook"
}

func (FishingHookType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

func Tick(e *Ent, target trace.Result) {
	fmt.Println("hook spawned")
}
