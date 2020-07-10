package reminder

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Psql struct {
	Pool *pgxpool.Pool
}

func (p *Psql) Insert(args InsertArgs) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*7)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	// Insert new reminder
	if _, err = conn.Exec(ctx, insertQuery, args.Id, args.GuildId, args.ChannelId, args.T, args.Name); err != nil {
		return err
	}

	return nil
}

func (p *Psql) ListFromGuild(args GuildListArgs) (*GuildListRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// Query by Guild ID
	rows, err := conn.Query(ctx, guildListQuery, args.GuildId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Iterate through results
	res := &GuildListRes{Data: []Reminder{}}
	for rows.Next() {
		var r Reminder

		err = rows.Scan(&r.Id, &r.T, &r.Name)
		if err != nil {
			return nil, err
		}

		res.Data = append(res.Data, r)
	}

	return res, nil
}

func (p *Psql) ListFromChannel(args ChannelListArgs) (*ChannelListRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// Query by Guild ID
	rows, err := conn.Query(ctx, channelListQuery, args.GuildId, args.ChannelId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Iterate through results
	res := &ChannelListRes{Data: []Reminder{}}
	for rows.Next() {
		var r Reminder

		err = rows.Scan(&r.Id, &r.T, &r.Name)
		if err != nil {
			return nil, err
		}

		res.Data = append(res.Data, r)
	}

	return res, nil
}

const (
	insertQuery      = `INSERT INTO reminders VALUES ($1, $2, $3, $4, $5)`
	guildListQuery   = `SELECT id, time, name FROM reminders WHERE guild_id = $1`
	channelListQuery = `SELECT id, time, name FROM reminders WHERE guild_id = $1 AND channel_id = $2`
)
