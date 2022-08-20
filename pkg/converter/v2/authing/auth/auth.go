package auth

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/auth"
)

func Ent2Grpc(row *ent.Auth) *npool.Auth {
	if row == nil {
		return nil
	}

	return &npool.Auth{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		RoleID:    row.RoleID.String(),
		UserID:    row.UserID.String(),
		Resource:  row.Resource,
		Method:    row.Method,
		CreatedAt: row.CreatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Auth) []*npool.Auth {
	infos := []*npool.Auth{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
