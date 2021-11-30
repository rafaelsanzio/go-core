package key

type Key struct {
	Name   string
	Secure bool
	Provider
}

type Provider string

var (
	ProviderStore  = Provider("store")
	ProviderEnvVar = Provider("env")
)

var (
	MongoDBName     = Key{Name: "db-name", Secure: false, Provider: ProviderStore}
	MongoDBPassword = Key{Name: "db-password", Secure: true, Provider: ProviderStore}
	MongoDBUsername = Key{Name: "db-user", Secure: false, Provider: ProviderStore}
	MongoURI        = Key{Name: "mongo-uri", Secure: false, Provider: ProviderStore}
	Region          = Key{Name: "region", Secure: false, Provider: ProviderStore}
)
