package kyc

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/kyc"
)

func trace(span trace1.Span, in *npool.KycReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("DocumentType.%v", index), in.GetDocumentType().String()),
		attribute.String(fmt.Sprintf("IDNumber.%v", index), in.GetIDNumber()),
		attribute.String(fmt.Sprintf("FrontImg.%v", index), in.GetFrontImg()),
		attribute.String(fmt.Sprintf("BackImg.%v", index), in.GetBackImg()),
		attribute.String(fmt.Sprintf("SelfieImg.%v", index), in.GetSelfieImg()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.KycReq) trace1.Span {
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
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.KycReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
