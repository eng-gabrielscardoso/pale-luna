package game

import (
	"github.com/eng-gabrielscardoso/pale-luna/internal/ai"
	"github.com/eng-gabrielscardoso/pale-luna/internal/config"
)

type State struct {
	PlayerName    string
	CurrentHour   int
	SessionCount  int
	PaleLunaAwake bool
	GameRunning   bool
	FirstTime     bool
	DebugMode     bool

	aiAgent *ai.AgentManager
	config  *config.Config
}

func NewGame(cfg *config.Config) *State {
	return &State{
		GameRunning:  true,
		FirstTime:    true,
		SessionCount: 0,
		config:       cfg,
		aiAgent:      ai.NewAgentManager(cfg),
	}
}

func (g *State) IsAIEnabled() bool {
	return g.aiAgent.IsAIAvailable()
}

func (g *State) GetAIStatus() map[string]interface{} {
	return g.aiAgent.GetStatus()
}
