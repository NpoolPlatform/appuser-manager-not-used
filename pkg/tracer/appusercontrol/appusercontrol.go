package appusercontrol

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusercontrol"
)

func trace(span trace1.Span, in *npool.AppUserControlReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.Bool(fmt.Sprintf("SigninVerifyByGoogleAuthentication.%v", index), in.GetSigninVerifyByGoogleAuthentication()),
		attribute.Bool(fmt.Sprintf("GoogleAuthenticationVerified.%v", index), in.GetGoogleAuthenticationVerified()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppUserControlReq) trace1.Span {
	return trace(span, in, 0)
}

func TraceConds(span trace1.Span, in *npool.Conds) trace1.Span {
	span.SetAttributes(
		attribute.String("ID.Op", in.GetID().GetOp()),
		attribute.String("ID.Val", in.GetID().GetValue()),
		attribute.String("AppID.Op", in.GetID().GetOp()),
		attribute.String("AppID.Val", in.GetID().GetValue()),
		attribute.String("UserID.Op", in.GetUserID().GetOp()),
		attribute.String("UserID.Val", in.GetUserID().GetValue()),
		attribute.String("SigninVerifyByGoogleAuthentication.Op", in.GetSigninVerifyByGoogleAuthentication().GetOp()),
		attribute.Bool("SigninVerifyByGoogleAuthentication.Val", in.GetSigninVerifyByGoogleAuthentication().GetValue()),
		attribute.String("GoogleAuthenticationVerified.Op", in.GetGoogleAuthenticationVerified().GetOp()),
		attribute.Bool("GoogleAuthenticationVerified.Val", in.GetGoogleAuthenticationVerified().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppUserControlReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
