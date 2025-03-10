package certcheck

import (
	"crypto/tls"
	_ "crypto/x509"
	_ "net"
	"testing"
	_ "time"
)

// TestCheckCertificate verifica a função CheckCertificate
func TestCheckCertificate(t *testing.T) {
	// Domínio de teste
	domain := "google.com"

	// Verifica o certificado
	info, err := CheckCertificate(domain)
	if err != nil {
		t.Fatalf("Erro ao verificar o certificado: %v", err)
	}

	// Verifica se as informações do certificado foram retornadas corretamente
	if info.Domain != domain {
		t.Errorf("Domínio esperado: %s, obtido: %s", domain, info.Domain)
	}
	if info.Issuer == "" {
		t.Error("Emitente do certificado não foi retornado")
	}
	if info.ExpiryDate.IsZero() {
		t.Error("Data de expiração do certificado não foi retornada")
	}
	if info.DaysRemaining < 0 {
		t.Error("Dias restantes não podem ser negativos")
	}
}

// TestCheckEncryptionAlgorithm verifica a função CheckEncryptionAlgorithm
func TestCheckEncryptionAlgorithm(t *testing.T) {
	// Conecta ao domínio e obtém o certificado
	conn, err := tls.Dial("tcp", "google.com:443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		t.Fatalf("Erro ao conectar ao domínio: %v", err)
	}
	defer conn.Close()

	// Obtém o certificado
	cert := conn.ConnectionState().PeerCertificates[0]

	// Verifica o algoritmo de criptografia
	result := CheckEncryptionAlgorithm(cert)
	if result == "" {
		t.Error("Resultado da verificação do algoritmo de criptografia não foi retornado")
	}
}
