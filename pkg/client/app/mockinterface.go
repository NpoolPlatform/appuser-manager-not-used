package app

import (
	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
)

type Repository interface {
	npool.AppMgrClient
}
