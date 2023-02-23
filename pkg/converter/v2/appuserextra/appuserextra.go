package appuserextra

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
)

func Ent2Grpc(row *ent.AppUserExtra) *npool.AppUserExtra {
	if row == nil {
		return nil
	}

	return &npool.AppUserExtra{
		PostalCode:    row.PostalCode,
		Avatar:        row.Avatar,
		Organization:  row.Organization,
		Birthday:      row.Birthday,
		ID:            row.ID.String(),
		AppID:         row.AppID.String(),
		Username:      row.Username,
		Gender:        row.Gender,
		LastName:      row.LastName,
		Age:           row.Age,
		UserID:        row.UserID.String(),
		FirstName:     row.FirstName,
		IDNumber:      row.IDNumber,
		AddressFields: row.AddressFields,
		ActionCredits: row.ActionCredits.String(),
	}
}

func Ent2GrpcMany(rows []*ent.AppUserExtra) []*npool.AppUserExtra {
	infos := []*npool.AppUserExtra{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
