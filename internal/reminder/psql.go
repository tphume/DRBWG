package reminder

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Psql struct {
	Pool *pgxpool.Pool
}

// This is duplicate of insert but whatever. I don't have the energy to refactor it
func (p *Psql) Set(args SetArgs) error {
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

func (p *Psql) Update(args *UpdateArgs) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	// Update by reminder id and guild id
	row := conn.QueryRow(ctx, updateQuery, args.T, args.Id, args.GuildId)
	if err := row.Scan(&args.Name); err != nil {
		return err
	}

	return nil
}

func (p *Psql) Del(args *DelArgs) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	// Delete by reminder id and guild id
	row := conn.QueryRow(ctx, deleteQuery, args.Id, args.GuildId)
	if err := row.Scan(&args.T, &args.Name); err != nil {
		return err
	}

	return nil
}

// Repo implementation for Pending type
type PsqlPending struct {
	Pool *pgxpool.Pool
}

func (p PsqlPending) GetPending(args GetPendingArgs) (*GetPendingRes, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Release()

	// Query by time
	rows, err := conn.Query(ctx, pendingQuery, args.Start, args.End)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// Iterate through results
	res := &GetPendingRes{Data: []Reminder{}}
	for rows.Next() {
		var r Reminder

		err = rows.Scan(&r.Id, &r.GuildId, &r.ChannelId, &r.Name, &r.T)
		if err != nil {
			return nil, err
		}

		res.Data = append(res.Data, r)
	}

	return res, nil
}

func (p PsqlPending) UpdateState(args StateArgs) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*12)
	defer cancel()

	// Get psql connection from pool
	conn, err := p.Pool.Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()

	// Update by reminder id and guild id
	if _, err = conn.Exec(ctx, stateQuery, args.Start, args.End); err != nil {
		return err
	}

	return nil
}

const (
	insertQuery      = `INSERT INTO reminders VALUES ($1, $2, $3, $4, $5)`
	guildListQuery   = `SELECT id, time, name FROM reminders WHERE guild_id = $1`
	channelListQuery = `SELECT id, time, name FROM reminders WHERE guild_id = $1 AND channel_id = $2`
	updateQuery      = `UPDATE reminders SET time = $1 WHERE id = $2 AND guild_id = $3 RETURNING name`
	deleteQuery      = `DELETE FROM reminders WHERE id = $1 AND guild_id = $2 RETURNING time, name`

	pendingQuery = `SELECT id, guild_id, channel_id, name, time FROM reminders WHERE done = FALSE AND time BETWEEN $1 AND $2`
	stateQuery   = `UPDATE reminders SET done = TRUE WHERE time BETWEEN $1 AND $2`
)
