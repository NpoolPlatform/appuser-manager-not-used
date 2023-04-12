package kyc

import (
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
)

func Ent2Grpc(row *ent.Kyc) *npool.Kyc {
	if row == nil {
		return nil
	}

	return &npool.Kyc{
		ID:           row.ID.String(),
		AppID:        row.AppID.String(),
		UserID:       row.UserID.String(),
		DocumentType: npool.KycDocumentType(npool.KycDocumentType_value[row.DocumentType]),
		IDNumber:     row.IDNumber,
		FrontImg:     row.FrontImg,
		BackImg:      row.BackImg,
		SelfieImg:    row.SelfieImg,
		EntityType:   npool.KycEntityType(npool.KycEntityType_value[row.EntityType]),
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
		ReviewID:     row.ReviewID.String(),
		State:        npool.KycState(npool.KycState_value[row.State]),
	}
}

func Ent2GrpcMany(rows []*ent.Kyc) []*npool.Kyc {
	infos := []*npool.Kyc{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
