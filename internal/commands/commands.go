package commands

import (
	"fmt"

	"github.com/carakacloud/ksema"
)

// CommandHandler handles all command execution
type CommandHandler struct {
	user *ksema.Ksema
}

// NewCommandHandler creates a new command handler
func NewCommandHandler(user *ksema.Ksema) *CommandHandler {
	return &CommandHandler{user: user}
}

// Execute executes the given command with arguments
func (h *CommandHandler) Execute(command string, args []string) error {
	switch command {
	case "help":
		return h.showHelp()
	case "encrypt":
		if h.user == nil {
			return fmt.Errorf("ksema connection required for encrypt command")
		}
		return h.handleEncrypt(args)
	case "decrypt":
		if h.user == nil {
			return fmt.Errorf("ksema connection required for decrypt command")
		}
		return h.handleDecrypt(args)
	case "sign":
		if h.user == nil {
			return fmt.Errorf("ksema connection required for sign command")
		}
		return h.handleSign(args)
	case "verify":
		if h.user == nil {
			return fmt.Errorf("ksema connection required for verify command")
		}
		return h.handleVerify(args)
	case "ping":
		if h.user == nil {
			return fmt.Errorf("ksema connection required for ping command")
		}
		return h.handlePing()
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

func (h *CommandHandler) showHelp() error {
	fmt.Println("Available commands:")
	fmt.Println("  encrypt <plaintext>                     - Encrypts the given plaintext")
	fmt.Println("  decrypt <ciphertext>                    - Decrypts the given ciphertext")
	fmt.Println("  sign <filename>                         - Signs the given file")
	fmt.Println("  verify <filename> <signature filename>  - Verifies the signature for the file")
	fmt.Println("  ping                                    - Pings the Ksema server")
	fmt.Println("  help                                    - Shows this help message")
	return nil
}

func (h *CommandHandler) handleEncrypt(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: encrypt <plaintext>")
	}
	plaintext := args[0]
	ciphertext, err := h.user.Encrypt([]byte(plaintext), "")
	if err != nil {
		return fmt.Errorf("encryption failed: %v", err)
	}
	fmt.Println(ciphertext)
	return nil
}

func (h *CommandHandler) handleDecrypt(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: decrypt <ciphertext>")
	}
	ciphertext := args[0]
	plaintext, err := h.user.Decrypt(ciphertext, "")
	if err != nil {
		return fmt.Errorf("decryption failed: %v", err)
	}
	fmt.Println(plaintext)
	return nil
}

func (h *CommandHandler) handleSign(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: sign <filename>")
	}
	filename := args[0]
	signatureFile, err := h.user.Sign(filename, "")
	if err != nil {
		return fmt.Errorf("signing failed: %v", err)
	}
	fmt.Println("Signature file created:", signatureFile)
	return nil
}

func (h *CommandHandler) handleVerify(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("usage: verify <filename> <signature filename>")
	}
	filename := args[0]
	signatureFile := args[1]
	err := h.user.Verify(filename, signatureFile, "")
	if err != nil {
		fmt.Println("Signature is invalid.")
		return fmt.Errorf("verification failed: %v", err)
	}
	fmt.Println("Signature is valid.")
	return nil
}

func (h *CommandHandler) handlePing() error {
	if err := h.user.Ping(); err != nil {
		return fmt.Errorf("ping failed: %v", err)
	}
	fmt.Println("Ping successful.")
	return nil
}
