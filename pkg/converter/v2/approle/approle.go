package approle

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
)

func Ent2Grpc(row *ent.AppRole) *npool.AppRole {
	if row == nil {
		return nil
	}

	return &npool.AppRole{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		CreatedBy:   row.CreatedBy.String(),
		Role:        row.Role,
		Description: row.Description,
		Default:     row.Default,
		Genesis:     row.Genesis,
	}
}

func Ent2GrpcMany(rows []*ent.AppRole) []*npool.AppRole {
	infos := []*npool.AppRole{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
