package model

import "time"

// Account is a Claude.ai account onboarded through the OAuth flow.
// Credential fields are stored encrypted with Config.CredentialEncryptionKey.
type Account struct {
	ID                 uint   `gorm:"primaryKey"`
	OrganizationID     string `gorm:"index"`
	Email              string
	EncryptedRefreshToken string
	EncryptedAccessToken  string
	AccessTokenExpiresAt  time.Time
	FingerprintTemplateID *uint
	Status             string // pending_delivery | delivered | revoked
	DeliveredAt        *time.Time
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

// FingerprintTemplate holds TLS/CC fingerprint parameters bound to accounts at onboarding time.
type FingerprintTemplate struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Params    string // JSON blob
	CreatedAt time.Time
}

// Operator is a platform admin user.
type Operator struct {
	ID           uint `gorm:"primaryKey"`
	Username     string `gorm:"uniqueIndex"`
	PasswordHash string
	CreatedAt    time.Time
}

// AuditLog records operator actions for compliance.
type AuditLog struct {
	ID         uint `gorm:"primaryKey"`
	OperatorID uint
	Action     string
	Detail     string
	CreatedAt  time.Time
}
