package appcontrol

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appcontrol"
)

func trace(span trace1.Span, in *npool.AppControlReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.Bool(fmt.Sprintf("KycEnable.%v", index), in.GetKycEnable()),
		attribute.Bool(fmt.Sprintf("SigninVerifyEnable.%v", index), in.GetSigninVerifyEnable()),
		attribute.Bool(fmt.Sprintf("InvitationCodeMust.%v", index), in.GetInvitationCodeMust()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppControlReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetAppID().GetOp()),
		attribute.String("AppID.Val", in.GetAppID().GetValue()),
		attribute.String("KycEnable.Op", in.GetKycEnable().GetOp()),
		attribute.Bool("KycEnable.Val", in.GetKycEnable().GetValue()),
		attribute.String("SigninVerifyEnable.Op", in.GetSigninVerifyEnable().GetOp()),
		attribute.Bool("SigninVerifyEnable.Val", in.GetSigninVerifyEnable().GetValue()),
		attribute.String("InvitationCodeMust.Op", in.GetInvitationCodeMust().GetOp()),
		attribute.Bool("InvitationCodeMust.Val", in.GetInvitationCodeMust().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppControlReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
