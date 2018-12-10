package wallet

// CreateWallet creates a new secure wallet with the given unique name
func CreateWallet(config WalletConfig, credential WalletCredential) error {
	channel := createWallet(config, credential)
	result := <-channel
	return result.Error
}
