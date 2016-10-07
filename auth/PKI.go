package auth

import (
	"fmt"
	"io/ioutil"
	"os"
)

// EnvDirectory is the env. variable name for the location of the directory where
// the PKI files are expected to be found.
const EnvDirectory = "TRIREME_PKI"

// DefaultPKIDirectory is the directory where the PEMs are mounted.
const DefaultPKIDirectory = "/var/trireme/"

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

// LoadPKI loads the PKI files based on the following directory:
// 1) Env Variable if set.
// 2) Default (/var/trireme/) if not set.
func LoadPKI() (*PKI, error) {
	dir := os.Getenv(EnvDirectory)
	if dir == "" {
		dir = DefaultPKIDirectory
	}
	return LoadPKIFromDir(dir)
}

// LoadPKIFromDir Create a new PKISecret from Kube Secret.
func LoadPKIFromDir(dir string) (*PKI, error) {
	keyPEM, err := ioutil.ReadFile(dir + KeyPEMFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read KeyPEMFile: %s", err)
	}
	certPEM, err := ioutil.ReadFile(dir + CertPEMFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read CertPEMFile: %s", err)
	}
	caCertPEM, err := ioutil.ReadFile(dir + CaCertPEMFile)
	if err != nil {
		return nil, fmt.Errorf("Couldn't read CaCertPEMFile %s", err)
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
