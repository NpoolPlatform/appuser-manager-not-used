package approle

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/approle"
)

func trace(span trace1.Span, in *npool.AppRoleReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("Role.%v", index), in.GetRole()),
		attribute.String(fmt.Sprintf("Description.%v", index), in.GetDescription()),
		attribute.String(fmt.Sprintf("CreatedBy.%v", index), in.GetCreatedBy()),
		attribute.Bool(fmt.Sprintf("Default.%v", index), in.GetDefault()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppRoleReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("Role.Op", in.GetRole().GetOp()),
		attribute.String("Role.Val", in.GetRole().GetValue()),
		attribute.String("CreatedBy.Op", in.GetCreatedBy().GetOp()),
		attribute.String("CreatedBy.Val", in.GetCreatedBy().GetValue()),
		attribute.String("Default.Op", in.GetDefault().GetOp()),
		attribute.Bool("Default.Val", in.GetDefault().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppRoleReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
