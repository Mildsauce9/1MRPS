package games

import (
	"context"

	repo "github.com/Mildsauce9/1MRPS/internal/adapters/postgresql/sqlc"
)

type GameProcessor interface {
	ListGames(ctx context.Context) ([]repo.Game, error)
}

type processor struct {
	// reference to the repo
	repo repo.Querier
}

func NewProcessor(repo repo.Querier) GameProcessor {
	return &processor{repo: repo}
}

func (p *processor) ListGames(ctx context.Context) ([]repo.Game, error) {
	games, err := p.repo.ListGames(ctx)
	return games, err

}
