package api

import (
	"github.com/NpoolPlatform/basal-manager/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/basal/mgr/v1/api"
)

func Ent2Grpc(row *ent.Api) *npool.API {
	if row == nil {
		return nil
	}

	return &npool.API{
		ID:              row.ID.String(),
		AppID:           row.AppID.String(),
		UserID:          row.UserID.String(),
		CoinTypeID:      row.CoinTypeID.String(),
		IOType:          npool.IOType(npool.IOType_value[row.IoType]),
		IOSubType:       npool.IOSubType(npool.IOSubType_value[row.IoSubType]),
		Amount:          row.Amount.String(),
		FromCoinTypeID:  row.FromCoinTypeID.String(),
		CoinUSDCurrency: row.CoinUsdCurrency.String(),
		IOExtra:         row.IoExtra,
		FromOldID:       row.FromOldID.String(),
	}
}

func Ent2GrpcMany(rows []*ent.Api) []*npool.API {
	infos := []*npool.API{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
