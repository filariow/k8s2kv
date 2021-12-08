package kv

import (
	"context"
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
	kvauth "github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
)

var ErrNotFound = errors.New("not found")

func UpdateSecret(ctx context.Context, vaultName, secret, value string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	p := keyvault.SecretSetParameters{Value: &value}
	v := buildKeyVaultUrl(vaultName)
	if _, err := c.SetSecret(ctx, v, secret, p); err != nil {
		return err
	}
	return nil
}

func GetSecret(ctx context.Context, vaultName, secret string) (*string, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	v := buildKeyVaultUrl(vaultName)
	s, err := c.GetSecret(ctx, v, secret, "")
	if err != nil {
		if s.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return s.Value, nil
}

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
