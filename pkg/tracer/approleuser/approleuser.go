package approleuser

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approleuser"
)

func trace(span trace1.Span, in *npool.AppRoleUserReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("RoleID.%v", index), in.GetRoleID()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppRoleUserReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("RoleID.Op", in.GetRoleID().GetOp()),
		attribute.String("RoleID.Val", in.GetRoleID().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppRoleUserReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
