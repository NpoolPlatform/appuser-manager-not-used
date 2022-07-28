package appusersecret

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appusersecret"
)

func trace(span trace1.Span, in *npool.AppUserSecretReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("PasswordHash.%v", index), in.GetPasswordHash()),
		attribute.String(fmt.Sprintf("Salt.%v", index), in.GetSalt()),
		attribute.String(fmt.Sprintf("GoogleSecret.%v", index), in.GetGoogleSecret()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppUserSecretReq) trace1.Span {
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
		attribute.String("PasswordHash.Op", in.GetPasswordHash().GetOp()),
		attribute.String("PasswordHash.Val", in.GetPasswordHash().GetValue()),
		attribute.String("Salt.Op", in.GetSalt().GetOp()),
		attribute.String("Salt.Val", in.GetSalt().GetValue()),
		attribute.String("GoogleSecret.Op", in.GetGoogleSecret().GetOp()),
		attribute.String("GoogleSecret.Val", in.GetGoogleSecret().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppUserSecretReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
