package indysdk_test

import (
	"errors"
	"testing"

	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk"
	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk/utils"
	"github.com/stretchr/testify/assert"
)

func testCreateWallet(t *testing.T) {
	assert := assert.New(t)

	config := indysdk.WalletConfig{
		ID: "wallet1",
	}
	credentials := indysdk.WalletCredential{
		Key: "wallet1",
	}
	assert.Equal(nil, indysdk.CreateWallet(config, credentials))
}

func testCreateWalletWithDuplicateID(t *testing.T) {
	assert := assert.New(t)

	config := indysdk.WalletConfig{
		ID: "wallet1",
	}
	credentials := indysdk.WalletCredential{
		Key: "wallet1",
	}

	err := errors.New(utils.GetIndyError(203))
	assert.Equal(err, indysdk.CreateWallet(config, credentials))
}

func TestWallet(t *testing.T) {
	testCreateWallet(t)
	testCreateWalletWithDuplicateID(t)
}
