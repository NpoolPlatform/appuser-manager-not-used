package banapp

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/banapp"
)

func trace(span trace1.Span, in *npool.BanAppReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("Message.%v", index), in.GetMessage()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.BanAppReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("Message.Op", in.GetMessage().GetOp()),
		attribute.String("Message.Val", in.GetMessage().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.BanAppReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
