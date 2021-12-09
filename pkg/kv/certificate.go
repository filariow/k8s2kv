package kv

import (
	"context"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/keyvault/keyvault"
)

func ImportCertificate(ctx context.Context, vaultName, certName, b64cert string) error {
	c, err := newClient()
	if err != nil {
		return err
	}

	p := keyvault.CertificateImportParameters{
		Base64EncodedCertificate: &b64cert,
	}
	u := buildKeyVaultUrl(vaultName)
	if _, err := c.ImportCertificate(ctx, u, certName, p); err != nil {
		return err
	}
	return nil
}

func GetCertificate(ctx context.Context, vaultName, certName string) (*[]byte, error) {
	c, err := newClient()
	if err != nil {
		return nil, err
	}

	u := buildKeyVaultUrl(vaultName)
	ce, err := c.GetCertificate(ctx, u, certName, "")
	if err != nil {
		if ce.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return ce.Cer, nil
}
