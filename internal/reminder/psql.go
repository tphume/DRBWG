package reminder

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Psql struct {
	Pool *pgxpool.Pool
}

func (p *Psql) InsertRepo(args InsertArgs) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	// Insert new reminder
	if _, err = conn.Exec(ctx, query, query, args.Id, args.GuildId, args.ChannelId, args.T, args.Name); err != nil {
		return err
	}

	return nil
}

const query = `INSERT INTO reminders VALUES ($1, $2, $3, $4, $5)`
