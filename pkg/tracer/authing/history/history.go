package history

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/authing/history"
)

func trace(span trace1.Span, in *npool.HistoryReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("Resource.%v", index), in.GetResource()),
		attribute.String(fmt.Sprintf("Method.%v", index), in.GetMethod()),
		attribute.Bool(fmt.Sprintf("Allowed.%v", index), in.GetAllowed()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.HistoryReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("Resource.Op", in.GetResource().GetOp()),
		attribute.String("Resource.Val", in.GetResource().GetValue()),
		attribute.String("Method.Op", in.GetMethod().GetOp()),
		attribute.String("Method.Val", in.GetMethod().GetValue()),
		attribute.String("Allowed.Op", in.GetAllowed().GetOp()),
		attribute.Bool("Allowed.Val", in.GetAllowed().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.HistoryReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
