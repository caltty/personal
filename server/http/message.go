package server

//LdapSettings is an option in the config file
type LdapSettings struct {
	URL          string `json:"url" validate:"notBlank"`
	BindUser     string `json:"bindUser" validate:"notBlank"`
	BindPassword string `json:"bindPassword" validate:"notBlank"`
	BaseDn       string `json:"baseDn" validate:"notBlank"`
}

//JwtAuthenticationSettings is an option in the config file
type JwtAuthenticationSettings struct {
	TokenDuration                  int                            `json:"tokenDuration" validate:"notBlank"`
	SecretkeyManagerType           string                         `json:"secretKeyManagerType" validate:"notBlank"`
	MemorySecretKeyManagerSettings MemorySecretKeyManagerSettings `json:"memorySecretKeyManagerSettings,omitempty" validate:"conditional:(SecretkeyManagerType=memory)"`
}

//MemorySecretKeyManagerSettings is an option in the config file
type MemorySecretKeyManagerSettings struct {
	CleanupInterval int `json:"cleanupInterval" validate:"notBlank"`
}

//Config is the data model of the config file
type Config struct {
	Ldap              LdapSettings              `json:"ldap" validate:"required"`
	JwtAuthentication JwtAuthenticationSettings `json:"jwtAuthentication" validate:"required"`
}

// LdapAttr struct
type LdapAttr struct {
	Dn       string   `json:"dn" validate:"notBlank"`
	AttrType string   `json:"attrType" validate:"notBlank"`
	AttrVals []string `json:"attrVals" validate:"notBlank"`
}

// Account struct
type Account struct {
	Username string `json:"username" validate"notBlank"`
	Password string `json:"password" validate:"notBlank"`
}
