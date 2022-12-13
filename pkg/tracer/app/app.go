package app

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/app"
)

func trace(span trace1.Span, in *npool.AppReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("Description.%v", index), in.GetDescription()),
		attribute.String(fmt.Sprintf("CreatedBy.%v", index), in.GetCreatedBy()),
		attribute.String(fmt.Sprintf("Name.%v", index), in.GetName()),
		attribute.String(fmt.Sprintf("Logo.%v", index), in.GetLogo()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("Description.Op", in.GetDescription().GetOp()),
		attribute.String("Description.Val", in.GetDescription().GetValue()),
		attribute.String("CreatedBy.Op", in.GetCreatedBy().GetOp()),
		attribute.String("CreatedBy.Val", in.GetCreatedBy().GetValue()),
		attribute.String("Name.Op", in.GetName().GetOp()),
		attribute.String("Name.Val", in.GetName().GetValue()),
		attribute.String("Logo.Op", in.GetLogo().GetOp()),
		attribute.String("Logo.Val", in.GetLogo().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
