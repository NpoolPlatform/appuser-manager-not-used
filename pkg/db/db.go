package db

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"

	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/app"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appcontrol"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approle"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/approleuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuser"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusercontrol"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserextra"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appusersecret"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/appuserthirdparty"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banapp"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/banappuser"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
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

func alterColumnNames(next schema.Applier) schema.Applier {
	return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
		tables := []string{
			app.Table,
			appcontrol.Table,
			approle.Table,
			approleuser.Table,
			appuser.Table,
			appusercontrol.Table,
			appuserextra.Table,
			appusersecret.Table,
			appuserthirdparty.Table,
			banapp.Table,
			banappuser.Table,
		}

		columns := [][]string{
			[]string{"create_at", "created_at"},
			[]string{"update_at", "updated_at"},
			[]string{"delete_at", "deleted_at"},
		}

	changeTableColumn:
		for _, table := range tables {
			for _, column := range columns {
				dstColumn := column[0]
				srcColumn := column[1]

				query, args := entsql.
					Select("datetime_precision").
					From(entsql.Table("`information_schema`.`columns`")).
					Where(
						entsql.And(
							entsql.EQ("table_name", table),
							entsql.EQ("column_name", dstColumn),
						),
					).
					Count().
					Query()

				rows := entsql.Rows{}
				if err := conn.Query(ctx, query, args, &rows); err != nil {
					return err
				}

				for rows.Next() {
					count := 0
					if err := rows.Scan(&count); err != nil {
						return err
					}
					if count > 0 {
						rows.Close()
						continue changeTableColumn
					}
				}
				rows.Close()

				cb := entsql.
					Column(dstColumn).
					Type(field.TypeInt.String())

				query, args = entsql.
					AlterTable(table).
					ChangeColumn(srcColumn, cb).
					Query()
				query = query + " unsigned"
				if err := conn.Exec(ctx, query, args, nil); err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func Init() error {
	cli, err := client()
	if err != nil {
		return err
	}
	err = cli.Schema.Create(
		context.Background(),
		schema.WithApplyHook(alterColumnNames),
	)
	if err != nil {
		return err
	}
	return nil
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
