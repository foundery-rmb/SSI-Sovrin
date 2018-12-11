package wallet

// StorageConfig represents Indy wallet storage config
type StorageConfig struct {
	Path string `json:"path"`
}

// Config represents Indy wallet config
type Config struct {
	ID            string        `json:"id"`
	StorageType   string        `json:"storage_type"`
	StorageConfig StorageConfig `json:"storage_config"`
}

// Credential represents Indy wallet credential config
type Credential struct {
	Key                   string `json:"key"`
	Rekey                 string `json:"rekey"`
	StorageCredentials    string `json:"storage_credentials"`
	KeyDerivationMethod   string `json:"key_derivation_method"`
	ReKeyDerivationMethod string `json:"rekey_derivation_method"`
}
