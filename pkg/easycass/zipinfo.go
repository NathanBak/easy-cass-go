package easycass

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
)

// zipinfo is a simple struct used to store information read from the secure
// connect bundle zip file
type zipinfo struct {
	hostname  string
	port      int
	tlsConfig *tls.Config
	keyspace  string
}

func fromProperties(hostname, port, keyspace, cert, key, caCerts string) (*zipinfo, error) {
	if hostname == "" {
		return nil, errors.New("hostname must be set")
	}

	portVal, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("bad port value : %v", err)
	}

	certPEMBlock, err := base64.StdEncoding.DecodeString(cert)
	if err != nil {
		return nil, fmt.Errorf("cert decode error : %v", err)
	}

	keyPemBlock, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return nil, fmt.Errorf("cert decode error : %v", err)
	}

	pemCerts, err := base64.StdEncoding.DecodeString(caCerts)
	if err != nil {
		return nil, fmt.Errorf("cert decode error : %v", err)
	}

	return newZipinfo(hostname, portVal, keyspace, certPEMBlock, keyPemBlock, pemCerts)
}

func newZipinfo(hostname string, port int, keyspace string, certPEMBlock, keyPemBlock, pemCerts []byte) (*zipinfo, error) {

	// Setup tlsconfig based on the cert, key and ca.crt file contents
	cert, _ := tls.X509KeyPair(certPEMBlock, keyPemBlock)
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(pemCerts)
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      caCertPool,
		ServerName:   hostname,
	}

	return &zipinfo{
		hostname:  hostname,
		port:      port,
		tlsConfig: tlsConfig,
		keyspace:  keyspace,
	}, nil
}
