package plguins

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/event"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

// Must implement the Plugin interface (dragonfly/plugin.go)
// Can implement some Handler interfaces (dragonfly/player/handler.go)
// Is included in the main.go
type CoordsPlugin struct{}

func (c CoordsPlugin) HandleCommandExecution(p *player.Player, ctx *event.Context, command cmd.Command, args []string) {
	if command.Name() == "coords" {
		p.ShowCoordinates()
	}
}

func (c CoordsPlugin) Init(server *dragonfly.Server) {
	cmd.Register(cmd.New("coords", "View your coordinates", nil))
}

func (c CoordsPlugin) GetName() string {
	return "Coords v0.0.1"
}
