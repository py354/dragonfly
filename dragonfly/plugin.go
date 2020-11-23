package dragonfly

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/entity"
	"github.com/df-mc/dragonfly/dragonfly/entity/damage"
	"github.com/df-mc/dragonfly/dragonfly/entity/healing"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/item"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
)

// Plugin can be used to create an implementation of Handler interface (player/plugins.go), without implementing
// all Plugin methods, but only with necessary
type Plugin interface {
	Init(*Server)
	GetName() string
}

// Handler - supporting structure for Plugin that implements the Handler interface (player/plugins.go)
type Handler struct {
	plugins []Plugin
}

func MakePlugin(plugin Plugin) *Handler {
	return &Handler{
		plugins: []Plugin{
			plugin,
		},
	}
}

func JoinPlugins(plugins []Plugin) *Handler {
	return &Handler{
		plugins: plugins,
	}
}

func (handler Handler) HandleMove(p *player.Player, ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.MoveHandler); ok {
			h.HandleMove(p, ctx, newPos, newYaw, newPitch)
		}
	}
}

func (handler Handler) HandleTeleport(p *player.Player, ctx *event.Context, pos mgl64.Vec3) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.TeleportHandler); ok {
			h.HandleTeleport(p, ctx, pos)
		}
	}
}

func (handler Handler) HandleChat(p *player.Player, ctx *event.Context, message *string) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ChatHandler); ok {
			h.HandleChat(p, ctx, message)
		}
	}
}

func (handler Handler) HandleFoodLoss(p *player.Player, ctx *event.Context, from, to int) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.FoodLossHandler); ok {
			h.HandleFoodLoss(p, ctx, from, to)
		}
	}
}

func (handler Handler) HandleHeal(p *player.Player, ctx *event.Context, health *float64, src healing.Source) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.HealHandler); ok {
			h.HandleHeal(p, ctx, health, src)
		}
	}
}

func (handler Handler) HandleHurt(p *player.Player, ctx *event.Context, damage *float64, src damage.Source) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.HurtHandler); ok {
			h.HandleHurt(p, ctx, damage, src)
		}
	}
}

func (handler Handler) HandleDeath(p *player.Player, src damage.Source) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.DeathHandler); ok {
			h.HandleDeath(p, src)
		}
	}
}

func (handler Handler) HandleRespawn(p *player.Player, pos *mgl64.Vec3) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.RespawnHandler); ok {
			h.HandleRespawn(p, pos)
		}
	}
}

func (handler Handler) HandleStartBreak(p *player.Player, ctx *event.Context, pos world.BlockPos) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.StartBreakHandler); ok {
			h.HandleStartBreak(p, ctx, pos)
		}
	}
}

func (handler Handler) HandleBlockBreak(p *player.Player, ctx *event.Context, pos world.BlockPos) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.BlockBreakHandler); ok {
			h.HandleBlockBreak(p, ctx, pos)
		}
	}
}

func (handler Handler) HandleBlockPlace(p *player.Player, ctx *event.Context, pos world.BlockPos, b world.Block) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.BlockPlaceHandler); ok {
			h.HandleBlockPlace(p, ctx, pos, b)
		}
	}
}

func (handler Handler) HandleBlockPick(p *player.Player, ctx *event.Context, pos world.BlockPos, b world.Block) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.BlockPickHandler); ok {
			h.HandleBlockPick(p, ctx, pos, b)
		}
	}
}

func (handler Handler) HandleItemUse(p *player.Player, ctx *event.Context) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ItemUseHandler); ok {
			h.HandleItemUse(p, ctx)
		}
	}
}

func (handler Handler) HandleItemUseOnBlock(p *player.Player, ctx *event.Context, pos world.BlockPos, face world.Face, clickPos mgl64.Vec3) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ItemUseOnBlockHandler); ok {
			h.HandleItemUseOnBlock(p, ctx, pos, face, clickPos)
		}
	}
}

func (handler Handler) HandleItemUseOnEntity(p *player.Player, ctx *event.Context, e world.Entity) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ItemUseOnEntityHandler); ok {
			h.HandleItemUseOnEntity(p, ctx, e)
		}
	}
}

func (handler Handler) HandleAttackEntity(p *player.Player, ctx *event.Context, e world.Entity) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.AttackEntityHandler); ok {
			h.HandleAttackEntity(p, ctx, e)
		}
	}
}

func (handler Handler) HandleItemDamage(p *player.Player, ctx *event.Context, i item.Stack, damage int) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ItemDamageHandler); ok {
			h.HandleItemDamage(p, ctx, i, damage)
		}
	}
}

func (handler Handler) HandleItemPickup(p *player.Player, ctx *event.Context, i item.Stack) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ItemPickupHandler); ok {
			h.HandleItemPickup(p, ctx, i)
		}
	}
}

func (handler Handler) HandleItemDrop(p *player.Player, ctx *event.Context, e *entity.Item) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.ItemDropHandler); ok {
			h.HandleItemDrop(p, ctx, e)
		}
	}
}

func (handler Handler) HandleTransfer(p *player.Player, ctx *event.Context, addr *net.UDPAddr) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.TransferHandler); ok {
			h.HandleTransfer(p, ctx, addr)
		}
	}
}

func (handler Handler) HandleCommandExecution(p *player.Player, ctx *event.Context, command cmd.Command, args []string) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.CommandExecutionHandler); ok {
			h.HandleCommandExecution(p, ctx, command, args)
		}
	}
}

func (handler Handler) HandleQuit(p *player.Player) {
	for _, plugin := range handler.plugins {
		if h, ok := plugin.(player.QuitHandler); ok {
			h.HandleQuit(p)
		}
	}
}
