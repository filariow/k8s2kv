package kv

import (
	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
)

func buildKeyVaultUrl(vaultName string) string {
	return "https://" + vaultName + ".vault.azure.net"
}

func newClient() (*keyvault.BaseClient, error) {
	a, err := kvauth.NewAuthorizerFromEnvironment()
	if err != nil {
		if a, err = kvauth.NewAuthorizerFromCLI(); err != nil {
			return nil, err
		}
	}

	c := keyvault.New()
	c.Authorizer = a
	return &c, err
}
