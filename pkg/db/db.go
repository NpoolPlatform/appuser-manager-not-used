package db

import (
	"context"
	"entgo.io/ent/dialect/sql/schema"
	"fmt"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/migrate"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"

	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"

	// ent policy runtime
	_ "github.com/NpoolPlatform/appuser-manager/pkg/db/ent/runtime"
)

func client() (*ent.Client, error) {
	conn, err := mysql.GetConn()
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init() error {
	cli, err := client()
	if err != nil {
		return err
	}

	return cli.Schema.Create(context.Background(), migrate.WithDropColumn(true), schema.WithDiffHook(renameColumnHook))
}

func Client() (*ent.Client, error) {
	return client()
}

func WithTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return err
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return fmt.Errorf("fail get client transaction: %v", err)
	}

	succ := false
	defer func() {
		if !succ {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail rollback: %v", err)
				return
			}
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %v", err)
	}

	succ = true
	return nil
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	cli, err := Client()
	if err != nil {
		return fmt.Errorf("fail get db client: %v", err)
	}

	if err := fn(ctx, cli); err != nil {
		return err
	}
	return nil
}

func renameColumnHook(next schema.Differ) schema.Differ {
	return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
		changes, err := next.Diff(current, desired)
		if err != nil {
			return nil, err
		}
		for _, c := range changes {
			m, ok := c.(*atlas.ModifyTable)

			if !ok {
				continue
			}

			changes := atlas.Changes(m.Changes)

			if i, j := changes.IndexDropColumn("create_at"), changes.IndexAddColumn("created_at"); i != -1 && j != -1 {
				changes = append(changes, &atlas.RenameColumn{
					From: changes[i].(*atlas.DropColumn).C,
					To:   changes[j].(*atlas.AddColumn).C,
				})
				changes.RemoveIndex(i, j)
			}

			if i, j := changes.IndexDropColumn("update_at"), changes.IndexAddColumn("updated_at"); i != -1 && j != -1 {
				changes = append(changes, &atlas.RenameColumn{
					From: changes[i].(*atlas.DropColumn).C,
					To:   changes[j].(*atlas.AddColumn).C,
				})
				changes.RemoveIndex(i, j)
			}

			if i, j := changes.IndexDropColumn("delete_at"), changes.IndexAddColumn("deleted_at"); i != -1 && j != -1 {
				changes = append(changes, &atlas.RenameColumn{
					From: changes[i].(*atlas.DropColumn).C,
					To:   changes[j].(*atlas.AddColumn).C,
				})
				changes.RemoveIndex(i, j)
			}

			m.Changes = changes
		}
		return changes, nil
	})
}
