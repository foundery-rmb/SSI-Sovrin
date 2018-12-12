package indysdk_test

import (
	"errors"
	"testing"

	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk"
	"github.com/sasiedu/SSI-Sovrin/indy-sdk-wrapper/golang/indysdk/utils"
	"github.com/stretchr/testify/assert"
)

var config = indysdk.WalletConfig{
	ID: "test_wallet",
}
var credentials = indysdk.WalletCredential{
	Key: "test_wallet",
}

func testCreateWallet(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(nil, indysdk.CreateWallet(config, credentials))
}

func testDeleteWallet(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(nil, indysdk.DeleteWallet(config, credentials))
}

func testCreateWalletWithDuplicateID(t *testing.T) {
	assert := assert.New(t)

	err := errors.New(utils.GetIndyError(203))
	assert.Equal(err, indysdk.CreateWallet(config, credentials))
}

func cleanUp() {
	indysdk.DeleteWallet(config, credentials)
}

func TestWallet(t *testing.T) {
	testCreateWallet(t)
	testCreateWalletWithDuplicateID(t)
	testDeleteWallet(t)
	cleanUp()
}
