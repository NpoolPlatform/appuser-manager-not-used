package appuserthirdparty

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
)

func Ent2Grpc(row *ent.AppUserThirdParty) *npool.AppUserThirdParty {
	if row == nil {
		return nil
	}

	return &npool.AppUserThirdParty{
		ID:                 row.ID.String(),
		AppID:              row.AppID.String(),
		UserID:             row.UserID.String(),
		ThirdPartyUserID:   row.ThirdPartyUserID,
		ThirdPartyID:       row.ThirdPartyID,
		ThirdPartyUsername: row.ThirdPartyUsername,
		ThirdPartyAvatar:   row.ThirdPartyAvatar,
	}
}

func Ent2GrpcMany(rows []*ent.AppUserThirdParty) []*npool.AppUserThirdParty {
	infos := []*npool.AppUserThirdParty{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
