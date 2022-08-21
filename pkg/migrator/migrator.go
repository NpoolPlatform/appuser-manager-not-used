package migrator

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	authingent "github.com/NpoolPlatform/authing-gateway/pkg/db/ent"
	authconstant "github.com/NpoolPlatform/authing-gateway/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	constant "github.com/NpoolPlatform/go-service-framework/pkg/mysql/const"
	sm "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
)

func Migrate(ctx context.Context) error {
	err := migrationAuthingGateway(ctx)
	if err != nil {
		return err
	}

	return setSigninVerifyTypeVal(ctx)
}

func setSigninVerifyTypeVal(ctx context.Context) error {
	cli, err := db.Client()
	if err != nil {
		return err
	}
	appUserControls, err := cli.AppUserControl.Query().All(ctx)
	if err != nil {
		return err
	}

	for _, val := range appUserControls {
		signinVerifyType := sm.SignMethodType_Email.String()
		if val.SigninVerifyByGoogleAuthentication {
			signinVerifyType = sm.SignMethodType_Google.String()
		}

		if _, err = cli.AppUserControl.
			UpdateOneID(val.ID).
			SetSigninVerifyType(signinVerifyType).
			Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

const (
	keyUsername = "username"
	keyPassword = "password"
	keyDBName   = "database_name"
	maxOpen     = 10
	maxIdle     = 10
)

func dsn(hostname string) (string, error) {
	username := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyUsername)
	password := config.GetStringValueWithNameSpace(constant.MysqlServiceName, keyPassword)
	dbname := config.GetStringValueWithNameSpace(hostname, keyDBName)

	svc, err := config.PeekService(constant.MysqlServiceName)
	if err != nil {
		logger.Sugar().Warnw("dsb", "error", err)
		return "", err
	}

	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true&interpolateParams=true",
		username, password,
		svc.Address,
		svc.Port,
		dbname,
	), nil
}

func open(hostname string) (conn *sql.DB, err error) {
	hdsn, err := dsn(hostname)
	if err != nil {
		return nil, err
	}

	conn, err = sql.Open("mysql", hdsn)
	if err != nil {
		return nil, err
	}

	// https://github.com/go-sql-driver/mysql
	// See "Important settings" section.

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(maxOpen)
	conn.SetMaxIdleConns(maxIdle)

	return conn, nil
}

func migrationAuthingGateway(ctx context.Context) (err error) {
	logger.Sugar().Infow("Migrate", "Start", "...")
	defer func() {
		logger.Sugar().Infow("Migrate", "Done", "...", "error", err)
	}()

	auth, err := open(authconstant.ServiceName)
	if err != nil {
		return err
	}

	authCli := authingent.NewClient(authingent.Driver(entsql.OpenDB(dialect.MySQL, auth)))

	cli, err := db.Client()
	if err != nil {
		return err
	}

	appAuths, err := authCli.
		AppAuth.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	appRoleAuths, err := authCli.
		AppRoleAuth.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	appUserAuths, err := authCli.
		AppUserAuth.
		Query().
		All(ctx)
	if err != nil {
		return err
	}

	auths := []*ent.Auth{}

	for _, val := range appAuths {
		auths = append(auths, &ent.Auth{
			AppID:     val.AppID,
			Resource:  val.Resource,
			Method:    val.Method,
			CreatedAt: val.CreateAt,
			UpdatedAt: val.UpdateAt,
		})
	}

	for _, val := range appRoleAuths {
		auths = append(auths, &ent.Auth{
			RoleID:    val.RoleID,
			Resource:  val.Resource,
			Method:    val.Method,
			CreatedAt: val.CreateAt,
			UpdatedAt: val.UpdateAt,
		})
	}

	for _, val := range appUserAuths {
		auths = append(auths, &ent.Auth{
			UserID:    val.UserID,
			Resource:  val.Resource,
			Method:    val.Method,
			CreatedAt: val.CreateAt,
			UpdatedAt: val.UpdateAt,
		})
	}

	tx, err := cli.Tx(ctx)
	if err != nil {
		return err
	}
	bulk := make([]*ent.AuthCreate, len(auths))
	for i, val := range auths {
		bulk[i] = tx.Auth.
			Create().
			SetAppID(val.AppID).
			SetRoleID(val.RoleID).
			SetUserID(val.UserID).
			SetResource(val.Resource).
			SetMethod(val.Method).
			SetCreatedAt(val.CreatedAt).
			SetUpdatedAt(val.UpdatedAt)
	}
	_, err = tx.Auth.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return err
	}

	return tx.Commit()
}
