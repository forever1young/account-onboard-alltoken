package service

// CredentialCipher encrypts/decrypts Account credential fields at rest using
// Config.CredentialEncryptionKey. This key must never be rotated once accounts
// exist in the database — rotating it makes existing credentials unrecoverable.
type CredentialCipher struct {
	key []byte
}

func NewCredentialCipher(key string) *CredentialCipher {
	return &CredentialCipher{key: []byte(key)}
}

func (c *CredentialCipher) Encrypt(plaintext string) (string, error) {
	panic("not implemented")
}

func (c *CredentialCipher) Decrypt(ciphertext string) (string, error) {
	panic("not implemented")
}
