package cmd

import (
	"crypto/tls"
	"log"

	"github.com/eduxcode/certguard/certcheck"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check [domínio]",
	Short: "Verifica o certificado SSL/TLS de um domínio.",
	Long:  `Verifica o certificado SSL/TLS de um domínio e exibe informações como validade, emitente e algoritmo de criptografia.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		domain := args[0]

		// Verifica o certificado
		info, err := certcheck.CheckCertificate(domain)
		if err != nil {
			log.Fatalf("Erro ao verificar o certificado: %v", err)
		}

        color.ForceOpenColor()
		// Exibe as informações com cores
		color.Cyan.Printf("🔒 Domínio: %s\n", info.Domain)
		color.Blue.Printf("📜 Emitido por: %s\n", info.Issuer)
		color.Yellow.Printf("📅 Data de expiração: %s\n", info.ExpiryDate.Format("2006-01-02"))
		color.Magenta.Printf("⏳ Dias restantes: %d\n", info.DaysRemaining)
		if info.IsValid {
			color.Green.Println("✅ Status: Válido")
		} else {
			color.Red.Println("❌ Status: Expirado")
		}

		// Verifica o algoritmo de criptografia
		conn, err := tls.Dial("tcp", domain+":443", &tls.Config{
			InsecureSkipVerify: true,
		})
		if err != nil {
			log.Fatalf("Erro ao conectar ao domínio: %v", err)
		}
		defer conn.Close()

		cert := conn.ConnectionState().PeerCertificates[0]
		algorithmStatus := certcheck.CheckEncryptionAlgorithm(cert)
		color.Cyan.Printf("🔐 Algoritmo de criptografia: %s\n", algorithmStatus)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
