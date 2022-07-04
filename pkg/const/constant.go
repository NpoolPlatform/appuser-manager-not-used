package constant

import (
	"time"
)

type DTMAction struct {
	Action string
	Revert string
}

const (
	SignupByMobile = "mobile"
	SignupByEmail  = "email"

	RecaptchaGoogleV3 = "google-recaptcha-v3"

	DBTimeout = 5 * time.Second

	GenesisRole = "genesis"

	GenesisAppName = "Genesis Dashboard"
	GenesisAppID   = "7203f5c0-7da9-11ec-a3ee-069013a3cb9a"

	ChurchAppName = "Church Dashboard"
	ChurchAppID   = "ab4d1208-7da9-11ec-a6ea-fb41bda845cd"

	CreateAppUserWithSecret       = "CreateAppUserWithSecret"
	CreateAppUserWithSecretRevert = "CreateAppUserWithSecretRevert"
)
