package games

import "context"

type GameProcessor interface {
	ListGames(ctx context.Context) error
}

type processor struct {
	// reference to the repo
}

func NewProcessor() GameProcessor {
	return &processor{}
}

func (p *processor) ListGames(ctx context.Context) error {
	return nil
}
