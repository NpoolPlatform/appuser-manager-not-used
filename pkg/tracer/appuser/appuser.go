package appuser

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuser"
)

func trace(span trace1.Span, in *npool.AppUserReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("EmailAddress.%v", index), in.GetEmailAddress()),
		attribute.String(fmt.Sprintf("PhoneNo.%v", index), in.GetPhoneNo()),
		attribute.String(fmt.Sprintf("ImportFromApp.%v", index), in.GetImportFromApp()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppUserReq) trace1.Span {
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
		attribute.String("PhoneNo.Op", in.GetPhoneNo().GetOp()),
		attribute.String("PhoneNo.Val", in.GetPhoneNo().GetValue()),
		attribute.String("ImportFromApp.Op", in.GetImportFromApp().GetOp()),
		attribute.String("ImportFromApp.Val", in.GetImportFromApp().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppUserReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
