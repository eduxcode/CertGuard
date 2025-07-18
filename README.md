# CertGuard

**CertGuard** é uma ferramenta em Go para verificar e monitorar certificados SSL/TLS. Ele fornece informações detalhadas sobre a validade, o emitente e o algoritmo de criptografia de um certificado, ajudando você a garantir que seus certificados estão seguros e atualizados.

---

## Funcionalidades

- Verifica a validade de certificados SSL/TLS.
- Exibe o emitente do certificado.
- Calcula os dias restantes até a expiração do certificado.
- Verifica o algoritmo de criptografia usado no certificado.
- Saída colorida no terminal para melhorar a legibilidade.

---

## Instalação

### Pré-requisitos

- Go 1.20 ou superior instalado.

### Passos para Instalação

1. Clone o repositório:
   ```bash
   git clone https://github.com/eduxcode/certguard.git
   cd certguard
   ```

2. Compile o projeto:
   ```bash
   go build -o certguard
   ```

3. (Opcional) Adicione o executável ao seu PATH para usar o CertGuard de qualquer lugar:
    - No Linux/macOS:
      ```bash
      export PATH=$PATH:$(pwd)
      ```
    - No Windows:
        - Adicione o diretório do executável à variável de ambiente `PATH`.

---

## Como Usar

### Verificar um Certificado

Para verificar o certificado SSL/TLS de um domínio, use o comando `check`:

```bash
./certguard check example.com
```

#### Saída Esperada:
```
🔒 Domínio: example.com
📜 Emitido por: DigiCert TLS RSA SHA256 2020 CA1
📅 Data de expiração: 2023-12-31
⏳ Dias restantes: 120
✅ Status: Válido
🔐 Algoritmo de criptografia: Algoritmo de criptografia seguro.
```

### Verificar Múltiplos Domínios

Você pode verificar vários domínios de uma vez executando o comando `check` para cada um:

```bash
./certguard check example.com
./certguard check google.com
./certguard check github.com
```

---

## Desenvolvimento

### Estrutura do Projeto

```
certguard/
├── cmd/
│   ├── root.go
│   └── check.go
├── certcheck/
│   ├── certcheck.go
│   └── certcheck_test.go
├── main.go
├── go.mod
├── go.sum
└── README.md
```

### Dependências

- [github.com/spf13/cobra](https://github.com/spf13/cobra): Para criar a interface de linha de comando (CLI).
- [github.com/gookit/color](https://github.com/gookit/color): Para adicionar cores à saída do terminal.

### Como Contribuir

1. Faça um fork do repositório.
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`).
3. Commit suas alterações (`git commit -m 'Adiciona nova feature'`).
4. Faça push para a branch (`git push origin feature/nova-feature`).
5. Abra um pull request.

---

## Licença

Este projeto está licenciado sob a licença MIT. Consulte o arquivo [LICENSE](https://github.com/eduxcode/CertGuard/blob/main/LICENSE) para mais detalhes.

---

## Autor

- **Davi Soares** ([GitHub](https://github.com/eduxcode))
