// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
	}
	// AppControlsColumns holds the columns for the "app_controls" table.
	AppControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "signup_methods", Type: field.TypeJSON, Nullable: true},
		{Name: "extern_signin_methods", Type: field.TypeJSON, Nullable: true},
		{Name: "recaptcha_method", Type: field.TypeString, Nullable: true, Default: "GoogleRecaptchaV3"},
		{Name: "kyc_enable", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "signin_verify_enable", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "invitation_code_must", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AppControlsTable holds the schema information for the "app_controls" table.
	AppControlsTable = &schema.Table{
		Name:       "app_controls",
		Columns:    AppControlsColumns,
		PrimaryKey: []*schema.Column{AppControlsColumns[0]},
	}
	// AppRolesColumns holds the columns for the "app_roles" table.
	AppRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "role", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "default", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "genesis", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AppRolesTable holds the schema information for the "app_roles" table.
	AppRolesTable = &schema.Table{
		Name:       "app_roles",
		Columns:    AppRolesColumns,
		PrimaryKey: []*schema.Column{AppRolesColumns[0]},
	}
	// AppRoleUsersColumns holds the columns for the "app_role_users" table.
	AppRoleUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "role_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppRoleUsersTable holds the schema information for the "app_role_users" table.
	AppRoleUsersTable = &schema.Table{
		Name:       "app_role_users",
		Columns:    AppRoleUsersColumns,
		PrimaryKey: []*schema.Column{AppRoleUsersColumns[0]},
	}
	// AppUsersColumns holds the columns for the "app_users" table.
	AppUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "email_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "phone_no", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "import_from_app", Type: field.TypeUUID, Nullable: true},
	}
	// AppUsersTable holds the schema information for the "app_users" table.
	AppUsersTable = &schema.Table{
		Name:       "app_users",
		Columns:    AppUsersColumns,
		PrimaryKey: []*schema.Column{AppUsersColumns[0]},
	}
	// AppUserControlsColumns holds the columns for the "app_user_controls" table.
	AppUserControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "signin_verify_by_google_authentication", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "google_authentication_verified", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "signin_verify_type", Type: field.TypeString, Nullable: true, Default: "Email"},
	}
	// AppUserControlsTable holds the schema information for the "app_user_controls" table.
	AppUserControlsTable = &schema.Table{
		Name:       "app_user_controls",
		Columns:    AppUserControlsColumns,
		PrimaryKey: []*schema.Column{AppUserControlsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appusercontrol_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserControlsColumns[4], AppUserControlsColumns[5]},
			},
		},
	}
	// AppUserExtrasColumns holds the columns for the "app_user_extras" table.
	AppUserExtrasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Default: ""},
		{Name: "first_name", Type: field.TypeString, Default: ""},
		{Name: "last_name", Type: field.TypeString, Default: ""},
		{Name: "address_fields", Type: field.TypeJSON},
		{Name: "gender", Type: field.TypeString, Default: ""},
		{Name: "postal_code", Type: field.TypeString, Default: ""},
		{Name: "age", Type: field.TypeUint32, Default: 0},
		{Name: "birthday", Type: field.TypeUint32, Default: 0},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "organization", Type: field.TypeString, Default: ""},
		{Name: "id_number", Type: field.TypeString, Default: ""},
	}
	// AppUserExtrasTable holds the schema information for the "app_user_extras" table.
	AppUserExtrasTable = &schema.Table{
		Name:       "app_user_extras",
		Columns:    AppUserExtrasColumns,
		PrimaryKey: []*schema.Column{AppUserExtrasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuserextra_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserExtrasColumns[4], AppUserExtrasColumns[5]},
			},
		},
	}
	// AppUserSecretsColumns holds the columns for the "app_user_secrets" table.
	AppUserSecretsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "google_secret", Type: field.TypeString, Default: ""},
	}
	// AppUserSecretsTable holds the schema information for the "app_user_secrets" table.
	AppUserSecretsTable = &schema.Table{
		Name:       "app_user_secrets",
		Columns:    AppUserSecretsColumns,
		PrimaryKey: []*schema.Column{AppUserSecretsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appusersecret_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserSecretsColumns[4], AppUserSecretsColumns[5]},
			},
		},
	}
	// AppUserThirdPartiesColumns holds the columns for the "app_user_third_parties" table.
	AppUserThirdPartiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "third_party_user_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "third_party_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "third_party_username", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "third_party_avatar", Type: field.TypeString, Nullable: true, Size: 1024, Default: ""},
	}
	// AppUserThirdPartiesTable holds the schema information for the "app_user_third_parties" table.
	AppUserThirdPartiesTable = &schema.Table{
		Name:       "app_user_third_parties",
		Columns:    AppUserThirdPartiesColumns,
		PrimaryKey: []*schema.Column{AppUserThirdPartiesColumns[0]},
	}
	// AuthsColumns holds the columns for the "auths" table.
	AuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "role_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "resource", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// AuthsTable holds the schema information for the "auths" table.
	AuthsTable = &schema.Table{
		Name:       "auths",
		Columns:    AuthsColumns,
		PrimaryKey: []*schema.Column{AuthsColumns[0]},
	}
	// AuthHistoriesColumns holds the columns for the "auth_histories" table.
	AuthHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "resource", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "allowed", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AuthHistoriesTable holds the schema information for the "auth_histories" table.
	AuthHistoriesTable = &schema.Table{
		Name:       "auth_histories",
		Columns:    AuthHistoriesColumns,
		PrimaryKey: []*schema.Column{AuthHistoriesColumns[0]},
	}
	// BanAppsColumns holds the columns for the "ban_apps" table.
	BanAppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Default: ""},
	}
	// BanAppsTable holds the schema information for the "ban_apps" table.
	BanAppsTable = &schema.Table{
		Name:       "ban_apps",
		Columns:    BanAppsColumns,
		PrimaryKey: []*schema.Column{BanAppsColumns[0]},
	}
	// BanAppUsersColumns holds the columns for the "ban_app_users" table.
	BanAppUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Default: ""},
	}
	// BanAppUsersTable holds the schema information for the "ban_app_users" table.
	BanAppUsersTable = &schema.Table{
		Name:       "ban_app_users",
		Columns:    BanAppUsersColumns,
		PrimaryKey: []*schema.Column{BanAppUsersColumns[0]},
	}
	// KycsColumns holds the columns for the "kycs" table.
	KycsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "document_type", Type: field.TypeString, Nullable: true, Default: "DefaultKycDocumentType"},
		{Name: "id_number", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "front_img", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "back_img", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "selfie_img", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "entity_type", Type: field.TypeString, Nullable: true, Default: "Individual"},
		{Name: "review_id", Type: field.TypeUUID, Nullable: true},
		{Name: "review_state", Type: field.TypeString, Nullable: true, Default: "DefaultReviewState"},
	}
	// KycsTable holds the schema information for the "kycs" table.
	KycsTable = &schema.Table{
		Name:       "kycs",
		Columns:    KycsColumns,
		PrimaryKey: []*schema.Column{KycsColumns[0]},
	}
	// LoginHistoriesColumns holds the columns for the "login_histories" table.
	LoginHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "client_ip", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "user_agent", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "location", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// LoginHistoriesTable holds the schema information for the "login_histories" table.
	LoginHistoriesTable = &schema.Table{
		Name:       "login_histories",
		Columns:    LoginHistoriesColumns,
		PrimaryKey: []*schema.Column{LoginHistoriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppsTable,
		AppControlsTable,
		AppRolesTable,
		AppRoleUsersTable,
		AppUsersTable,
		AppUserControlsTable,
		AppUserExtrasTable,
		AppUserSecretsTable,
		AppUserThirdPartiesTable,
		AuthsTable,
		AuthHistoriesTable,
		BanAppsTable,
		BanAppUsersTable,
		KycsTable,
		LoginHistoriesTable,
	}
)

func init() {
}
