package teams

import (
	"context"

	"github.com/keybase/client/go/libkb"
	"github.com/keybase/client/go/protocol/keybase1"
)

func HandleRotateRequest(ctx context.Context, g *libkb.GlobalContext, teamID keybase1.TeamID, generation keybase1.PerTeamKeyGeneration) error {
	team, err := GetForTeamManagement(ctx, g, teamID)
	if err != nil {
		return err
	}

	if team.Generation() > generation {
		g.Log.Debug("current team generation %d > team.clkr generation %d, not rotating", team.Generation(), generation)
		return nil
	}

	g.Log.Debug("rotating team %s (%s)", team.Name, teamID)
	if err := team.Rotate(ctx); err != nil {
		g.Log.Debug("rotating team %s (%s) error: %s", team.Name, teamID, err)
		return err
	}

	g.Log.Debug("sucess rotating team %s (%s)", team.Name, teamID)
	return nil
}
