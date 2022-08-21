package history

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/login/history"
)

func trace(span trace1.Span, in *npool.HistoryReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("ClientIP.%v", index), in.GetClientIP()),
		attribute.String(fmt.Sprintf("UserAgent.%v", index), in.GetUserAgent()),
		attribute.String(fmt.Sprintf("Location.%v", index), in.GetLocation()),
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
		attribute.String("ClientIP.Op", in.GetClientIP().GetOp()),
		attribute.String("ClientIP.Val", in.GetClientIP().GetValue()),
		attribute.String("UserAgent.Op", in.GetUserAgent().GetOp()),
		attribute.String("UserAgent.Val", in.GetUserAgent().GetValue()),
		attribute.String("Location.Op", in.GetLocation().GetOp()),
		attribute.String("Location.Val", in.GetLocation().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.HistoryReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
