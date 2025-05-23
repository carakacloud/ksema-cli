# ksema-cli

A command-line tool for interacting with KSEMA.

## Features

- **Encryption**: Securely encrypt plaintext data
- **Decryption**: Decrypt previously encrypted data
- **Digital Signatures**: Sign files to verify their authenticity
- **Signature Verification**: Verify the authenticity of signed files
- **KSEMA Server Health Check**: Ping the KSEMA server to check its status

## Installation

Clone the repository and build the binary:

```sh
git clone https://github.com/carakacloud/ksema-cli.git
cd ksema-cli
make build
```

## Usage

For help, run:

### Available Commands

```md
>> ksema-cli help
  encrypt <plaintext>                     - Encrypts the given plaintext
  decrypt <ciphertext>                    - Decrypts the given ciphertext
  sign <filename>                         - Signs the given file
  verify <filename> <signature filename>  - Verifies the signature for the file
  ping                                    - Pings the Ksema server
  help                                    - Shows this help message
```