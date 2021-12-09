package kv

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
)

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
