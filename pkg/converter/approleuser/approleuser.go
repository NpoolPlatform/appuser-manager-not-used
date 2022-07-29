package approleuser

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
)

func Ent2Grpc(row *ent.AppRoleUser) *npool.AppRoleUser {
	if row == nil {
		return nil
	}

	return &npool.AppRoleUser{
		AppID:  row.AppID.String(),
		RoleID: row.RoleID.String(),
		UserID: row.UserID.String(),
		ID:     row.ID.String(),
	}
}

func Ent2GrpcMany(rows []*ent.AppRoleUser) []*npool.AppRoleUser {
	infos := []*npool.AppRoleUser{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
