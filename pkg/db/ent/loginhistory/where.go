// Code generated by ent, DO NOT EDIT.

package loginhistory

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/appuser-manager/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ClientIP applies equality check predicate on the "client_ip" field. It's identical to ClientIPEQ.
func ClientIP(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientIP), v))
	})
}

// UserAgent applies equality check predicate on the "user_agent" field. It's identical to UserAgentEQ.
func UserAgent(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserAgent), v))
	})
}

// Location applies equality check predicate on the "location" field. It's identical to LocationEQ.
func Location(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// AppIDIsNil applies the IsNil predicate on the "app_id" field.
func AppIDIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppID)))
	})
}

// AppIDNotNil applies the NotNil predicate on the "app_id" field.
func AppIDNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppID)))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// ClientIPEQ applies the EQ predicate on the "client_ip" field.
func ClientIPEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldClientIP), v))
	})
}

// ClientIPNEQ applies the NEQ predicate on the "client_ip" field.
func ClientIPNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldClientIP), v))
	})
}

// ClientIPIn applies the In predicate on the "client_ip" field.
func ClientIPIn(vs ...string) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldClientIP), v...))
	})
}

// ClientIPNotIn applies the NotIn predicate on the "client_ip" field.
func ClientIPNotIn(vs ...string) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldClientIP), v...))
	})
}

// ClientIPGT applies the GT predicate on the "client_ip" field.
func ClientIPGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldClientIP), v))
	})
}

// ClientIPGTE applies the GTE predicate on the "client_ip" field.
func ClientIPGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldClientIP), v))
	})
}

// ClientIPLT applies the LT predicate on the "client_ip" field.
func ClientIPLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldClientIP), v))
	})
}

// ClientIPLTE applies the LTE predicate on the "client_ip" field.
func ClientIPLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldClientIP), v))
	})
}

// ClientIPContains applies the Contains predicate on the "client_ip" field.
func ClientIPContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldClientIP), v))
	})
}

// ClientIPHasPrefix applies the HasPrefix predicate on the "client_ip" field.
func ClientIPHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldClientIP), v))
	})
}

// ClientIPHasSuffix applies the HasSuffix predicate on the "client_ip" field.
func ClientIPHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldClientIP), v))
	})
}

// ClientIPIsNil applies the IsNil predicate on the "client_ip" field.
func ClientIPIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldClientIP)))
	})
}

// ClientIPNotNil applies the NotNil predicate on the "client_ip" field.
func ClientIPNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldClientIP)))
	})
}

// ClientIPEqualFold applies the EqualFold predicate on the "client_ip" field.
func ClientIPEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldClientIP), v))
	})
}

// ClientIPContainsFold applies the ContainsFold predicate on the "client_ip" field.
func ClientIPContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldClientIP), v))
	})
}

// UserAgentEQ applies the EQ predicate on the "user_agent" field.
func UserAgentEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserAgent), v))
	})
}

// UserAgentNEQ applies the NEQ predicate on the "user_agent" field.
func UserAgentNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserAgent), v))
	})
}

// UserAgentIn applies the In predicate on the "user_agent" field.
func UserAgentIn(vs ...string) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserAgent), v...))
	})
}

// UserAgentNotIn applies the NotIn predicate on the "user_agent" field.
func UserAgentNotIn(vs ...string) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserAgent), v...))
	})
}

// UserAgentGT applies the GT predicate on the "user_agent" field.
func UserAgentGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserAgent), v))
	})
}

// UserAgentGTE applies the GTE predicate on the "user_agent" field.
func UserAgentGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserAgent), v))
	})
}

// UserAgentLT applies the LT predicate on the "user_agent" field.
func UserAgentLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserAgent), v))
	})
}

// UserAgentLTE applies the LTE predicate on the "user_agent" field.
func UserAgentLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserAgent), v))
	})
}

// UserAgentContains applies the Contains predicate on the "user_agent" field.
func UserAgentContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserAgent), v))
	})
}

// UserAgentHasPrefix applies the HasPrefix predicate on the "user_agent" field.
func UserAgentHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserAgent), v))
	})
}

// UserAgentHasSuffix applies the HasSuffix predicate on the "user_agent" field.
func UserAgentHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserAgent), v))
	})
}

// UserAgentIsNil applies the IsNil predicate on the "user_agent" field.
func UserAgentIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserAgent)))
	})
}

// UserAgentNotNil applies the NotNil predicate on the "user_agent" field.
func UserAgentNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserAgent)))
	})
}

// UserAgentEqualFold applies the EqualFold predicate on the "user_agent" field.
func UserAgentEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserAgent), v))
	})
}

// UserAgentContainsFold applies the ContainsFold predicate on the "user_agent" field.
func UserAgentContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserAgent), v))
	})
}

// LocationEQ applies the EQ predicate on the "location" field.
func LocationEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocation), v))
	})
}

// LocationNEQ applies the NEQ predicate on the "location" field.
func LocationNEQ(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocation), v))
	})
}

// LocationIn applies the In predicate on the "location" field.
func LocationIn(vs ...string) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocation), v...))
	})
}

// LocationNotIn applies the NotIn predicate on the "location" field.
func LocationNotIn(vs ...string) predicate.LoginHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocation), v...))
	})
}

// LocationGT applies the GT predicate on the "location" field.
func LocationGT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocation), v))
	})
}

// LocationGTE applies the GTE predicate on the "location" field.
func LocationGTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocation), v))
	})
}

// LocationLT applies the LT predicate on the "location" field.
func LocationLT(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocation), v))
	})
}

// LocationLTE applies the LTE predicate on the "location" field.
func LocationLTE(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocation), v))
	})
}

// LocationContains applies the Contains predicate on the "location" field.
func LocationContains(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLocation), v))
	})
}

// LocationHasPrefix applies the HasPrefix predicate on the "location" field.
func LocationHasPrefix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLocation), v))
	})
}

// LocationHasSuffix applies the HasSuffix predicate on the "location" field.
func LocationHasSuffix(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLocation), v))
	})
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLocation)))
	})
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLocation)))
	})
}

// LocationEqualFold applies the EqualFold predicate on the "location" field.
func LocationEqualFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLocation), v))
	})
}

// LocationContainsFold applies the ContainsFold predicate on the "location" field.
func LocationContainsFold(v string) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLocation), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.LoginHistory) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.LoginHistory) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.LoginHistory) predicate.LoginHistory {
	return predicate.LoginHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}
