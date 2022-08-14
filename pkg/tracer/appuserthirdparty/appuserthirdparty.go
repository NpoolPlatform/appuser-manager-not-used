package appuserthirdparty

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserthirdparty"
)

func trace(span trace1.Span, in *npool.AppUserThirdPartyReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.String(fmt.Sprintf("ThirdPartyUserID.%v", index), in.GetThirdPartyUserID()),
		attribute.String(fmt.Sprintf("ThirdPartyID.%v", index), in.GetThirdPartyID()),
		attribute.String(fmt.Sprintf("ThirdPartyUsername.%v", index), in.GetThirdPartyUsername()),
		attribute.String(fmt.Sprintf("ThirdPartyAvatar.%v", index), in.GetThirdPartyAvatar()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppUserThirdPartyReq) trace1.Span {
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
		attribute.String("ThirdPartyUserID.Op", in.GetThirdPartyUserID().GetOp()),
		attribute.String("ThirdPartyUserID.Val", in.GetThirdPartyUserID().GetValue()),
		attribute.String("ThirdPartyID.Op", in.GetThirdPartyID().GetOp()),
		attribute.String("ThirdPartyID.Val", in.GetThirdPartyID().GetValue()),
		attribute.String("ThirdPartyUsername.Op", in.GetThirdPartyUsername().GetOp()),
		attribute.String("ThirdPartyUsername.Val", in.GetThirdPartyUsername().GetValue()),
		attribute.String("ThirdPartyAvatar.Op", in.GetThirdPartyAvatar().GetOp()),
		attribute.String("ThirdPartyAvatar.Val", in.GetThirdPartyAvatar().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppUserThirdPartyReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
