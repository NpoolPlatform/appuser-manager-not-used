package subscriber

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/subscriber"
)

func trace(span trace1.Span, in *npool.SubscriberReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("EmailAddress.%v", index), in.GetEmailAddress()),
		attribute.Bool(fmt.Sprintf("Registered.%v", index), in.GetRegistered()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.SubscriberReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("EmailAddress.Op", in.GetEmailAddress().GetOp()),
		attribute.String("EmailAddress.Val", in.GetEmailAddress().GetValue()),
		attribute.String("Registered.Op", in.GetRegistered().GetOp()),
		attribute.Bool("Registered.Val", in.GetRegistered().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.SubscriberReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
