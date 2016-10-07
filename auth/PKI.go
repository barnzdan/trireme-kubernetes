package auth

import (
	"fmt"
	"io/ioutil"
)

// SecretDirectory is the directory where the PEMs are mounted.
const SecretDirectory = "/var/trireme/"

// KeyPEMFile is the name of the KeyPEMFile in the SecretDirectory directory.
const KeyPEMFile = "key.pem"

// CertPEMFile is the name of the CertPEsMFile in the SecretDirectory directory.
const CertPEMFile = "cert.pem"

// CaCertPEMFile is the name of the CaCertPEMFile in the SecretDirectory directory.
const CaCertPEMFile = "ca.crt"

// A PKI is used to
type PKI struct {
	KeyPEM    []byte
	CertPEM   []byte
	CaCertPEM []byte
}

// LoadPKIFromKubeSecret Create a new PKISecret from Kube Secret.
func LoadPKIFromKubeSecret() (*PKI, error) {
	keyPEM, err := ioutil.ReadFile(SecretDirectory + KeyPEMFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read KeyPEMFile")
	}
	certPEM, err := ioutil.ReadFile(SecretDirectory + CertPEMFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read CertPEMFile")
	}
	caCertPEM, err := ioutil.ReadFile(SecretDirectory + CaCertPEMFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read CaCertPEMFile")
	}

	return &PKI{
		KeyPEM:    keyPEM,
		CertPEM:   certPEM,
		CaCertPEM: caCertPEM,
	}, nil
}

// PublishPKI Publishes on the node annotation the PublicKey to use for this node
func (p *PKI) PublishPKI() {

}
