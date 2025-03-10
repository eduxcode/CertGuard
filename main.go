package main

import (
	"crypto/tls" // Importação necessária para tls.Dial
	"fmt"
	"log"
	"github.com/eduxcode/certguard/certcheck"
)

func main() {
	// Domínio a ser verificado
	domain := "gogle.com"

	// Verifica o certificado
	info, err := certcheck.CheckCertificate(domain)
	if err != nil {
		log.Fatalf("Erro ao verificar o certificado: %v", err)
	}

	// Exibe as informações
	fmt.Printf("Domínio: %s\n", info.Domain)
	fmt.Printf("Emitido por: %s\n", info.Issuer)
	fmt.Printf("Data de expiração: %s\n", info.ExpiryDate.Format("2006-01-02"))
	fmt.Printf("Dias restantes: %d\n", info.DaysRemaining)
	if info.IsValid {
		fmt.Println("Status: Válido")
	} else {
		fmt.Println("Status: Expirado")
	}

	// Verifica o algoritmo de criptografia
	conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
		InsecureSkipVerify: true, // Ignora erros de certificado inválido
	})
	if err != nil {
		log.Fatalf("Erro ao conectar ao domínio: %v", err)
	}
	defer conn.Close()

	// Obtém o certificado
	cert := conn.ConnectionState().PeerCertificates[0]

	// Verifica o algoritmo de criptografia
	algorithmStatus := certcheck.CheckEncryptionAlgorithm(cert)
	fmt.Println(algorithmStatus)
}
