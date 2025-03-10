package certcheck

import (
	"crypto/tls" // Para tls.Dial
	"crypto/x509" // Para analisar o certificado
	"fmt"
	_"net" // Importação adicionada para operações de rede
	"time" // Para calcular a validade do certificado
)

// CertificateInfo contém informações sobre o certificado
type CertificateInfo struct {
	Domain       string
	Issuer       string
	ExpiryDate   time.Time
	DaysRemaining int
	IsValid      bool
}

// CheckCertificate verifica o certificado de um domínio
func CheckCertificate(domain string) (*CertificateInfo, error) {
	// Conecta ao domínio e obtém o certificado
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true, // Ignora erros de certificado inválido
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao domínio: %v", err)
	}
	defer conn.Close()

	// Obtém o certificado
	cert := conn.ConnectionState().PeerCertificates[0]

	// Verifica a validade do certificado
	now := time.Now()
	expiryDate := cert.NotAfter
	daysRemaining := int(expiryDate.Sub(now).Hours() / 24)

	// Retorna as informações do certificado
	info := &CertificateInfo{
		Domain:       domain,
		Issuer:       cert.Issuer.CommonName,
		ExpiryDate:   expiryDate,
		DaysRemaining: daysRemaining,
		IsValid:      now.Before(expiryDate),
	}

	return info, nil
}

// CheckEncryptionAlgorithm verifica o algoritmo de criptografia do certificado
func CheckEncryptionAlgorithm(cert *x509.Certificate) string {
	switch cert.SignatureAlgorithm {
	case x509.SHA1WithRSA, x509.DSAWithSHA1, x509.ECDSAWithSHA1:
		return "AVISO: O certificado usa SHA-1, que é considerado inseguro."
	default:
		return "Algoritmo de criptografia seguro."
	}
}
