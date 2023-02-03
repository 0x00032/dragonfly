package item

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/world"
)

// FishingRod is an item used to catch fish in water.
type FishingRod struct{}

// DurabilityInfo Fishing Rods in minecraft bedrock have a 384 durability and return nothing when broken.
func (FishingRod) DurabilityInfo() DurabilityInfo {
	return DurabilityInfo{
		MaxDurability: 384,
		BrokenItem:    simpleItem(Stack{}),
	}
}

// MaxCount Fishing Rods are Non-Stackable items.
func (FishingRod) MaxCount() int {
	return 1
}

// Use upon using a fishing rod we lower durability by one and spawn a FishingHook entity.
func (FishingRod) Use(w *world.World, u User, ctx *UseContext) bool {

	// Debug
	fmt.Println(w)
	fmt.Println(u)

	ctx.DamageItem(1)
	return true
}

// EncodeItem For minecraft's string item system.
func (FishingRod) EncodeItem() (name string, meta int16) {
	return "minecraft:fishing_rod", 0
}
