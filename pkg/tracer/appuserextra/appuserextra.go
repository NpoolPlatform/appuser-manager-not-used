package appuserextra

import (
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	trace1 "go.opentelemetry.io/otel/trace"

	npool "github.com/NpoolPlatform/message/npool/appuser/mgr/v2/appuserextra"
)

func trace(span trace1.Span, in *npool.AppUserExtraReq, index int) trace1.Span {
	span.SetAttributes(
		attribute.String(fmt.Sprintf("ID.%v", index), in.GetID()),
		attribute.String(fmt.Sprintf("AppID.%v", index), in.GetAppID()),
		attribute.String(fmt.Sprintf("UserID.%v", index), in.GetUserID()),
		attribute.StringSlice(fmt.Sprintf("AddressFields.%v", index), in.GetAddressFields()),
		attribute.String(fmt.Sprintf("Username.%v", index), in.GetUsername()),
		attribute.Int(fmt.Sprintf("Age.%v", index), int(in.GetAge())),
		attribute.String(fmt.Sprintf("Avatar.%v", index), in.GetAvatar()),
		attribute.Int(fmt.Sprintf("Birthday.%v", index), int(in.GetBirthday())),
		attribute.String(fmt.Sprintf("FirstName.%v", index), in.GetFirstName()),
		attribute.String(fmt.Sprintf("Gender.%v", index), in.GetGender()),
		attribute.String(fmt.Sprintf("IDNumber.%v", index), in.GetIDNumber()),
		attribute.String(fmt.Sprintf("LastName.%v", index), in.GetLastName()),
		attribute.String(fmt.Sprintf("Organization.%v", index), in.GetOrganization()),
		attribute.String(fmt.Sprintf("PostalCode.%v", index), in.GetPostalCode()),
	)
	return span
}

func Trace(span trace1.Span, in *npool.AppUserExtraReq) trace1.Span {
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
		attribute.String("Username.Op", in.GetUsername().GetOp()),
		attribute.String("Username.Val", in.GetUsername().GetValue()),
		attribute.String("Age.Op", in.GetAge().GetOp()),
		attribute.Int("Age.Val", int(in.GetAge().GetValue())),
		attribute.String("Avatar.Op", in.GetAvatar().GetOp()),
		attribute.String("Avatar.Val", in.GetAvatar().GetValue()),
		attribute.String("Birthday.Op", in.GetBirthday().GetOp()),
		attribute.Int("Birthday.Val", int(in.GetBirthday().GetValue())),
		attribute.String("FirstName.Op", in.GetFirstName().GetOp()),
		attribute.String("FirstName.Val", in.GetFirstName().GetValue()),
		attribute.String("Gender.Op", in.GetGender().GetOp()),
		attribute.String("Gender.Val", in.GetGender().GetValue()),
		attribute.String("IDNumber.Op", in.GetIDNumber().GetOp()),
		attribute.String("IDNumber.Val", in.GetIDNumber().GetValue()),
		attribute.String("LastName.Op", in.GetLastName().GetOp()),
		attribute.String("LastName.Val", in.GetLastName().GetValue()),
		attribute.String("Organization.Op", in.GetOrganization().GetOp()),
		attribute.String("Organization.Val", in.GetOrganization().GetValue()),
		attribute.String("PostalCode.Op", in.GetPostalCode().GetOp()),
		attribute.String("PostalCode.Val", in.GetPostalCode().GetValue()),
	)
	return span
}

func TraceMany(span trace1.Span, infos []*npool.AppUserExtraReq) trace1.Span {
	for index, info := range infos {
		span = trace(span, info, index)
	}
	return span
}
