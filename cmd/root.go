package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "certguard",
	Short: "CertGuard é uma ferramenta para verificar certificados SSL/TLS.",
	Long: `CertGuard é uma ferramenta em Go que analisa e monitora certificados SSL/TLS,
fornecendo informações sobre validade, algoritmos de criptografia e muito mais.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
