package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

const encodedHashFormat = "$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s"

func defaultArgon2Params() *argon2Params {
	return &argon2Params{
		Memory:      64 * 1024, // 64 MB
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
}

func GenerateMFACode() (string, error) {
	var max uint32 = 1000000 // 6 digits
	var num uint32

	// Read 4 random bytes
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// Convert bytes to uint32
	num = (uint32(b[0]) << 24) | (uint32(b[1]) << 16) | (uint32(b[2]) << 8) | uint32(b[3])
	code := num % max

	return fmt.Sprintf("%06d", code), nil
}

type argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func IsValidPassword(pwd string) bool {
	if len(pwd) < 10 {
		return false
	}
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(pwd)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(pwd)
	hasSpecial := regexp.MustCompile(`[^a-zA-Z0-9]`).MatchString(pwd)

	return hasUpper && hasNumber && hasSpecial
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func HashPassword(password string) (encodedHash string, err error) {
	return hashPassword(password, defaultArgon2Params())
}

func hashPassword(password string, p *argon2Params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Base64 encode salt and hash for storage
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Format: $argon2id$v=19$m=65536,t=3,p=2$<salt>$<hash>
	encodedHash = fmt.Sprintf(encodedHashFormat, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)
	return encodedHash, nil
}

func VerifyPassword(password, encodedHash string) (bool, error) {
	var salt, hash []byte

	parts := split(encodedHash, "$")
	// parts[0] = "" (empty)
	// parts[1] = argon2id
	// parts[2] = v=19
	// parts[3] = m=65536,t=3,p=2
	// parts[4] = salt (base64)
	// parts[5] = hash (base64)

	if len(parts) != 6 {
		return false, fmt.Errorf("invalid hash format")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false, fmt.Errorf("unable to scanf version: %w", err)
	}
	if version != 19 {
		return false, fmt.Errorf("unsupported version %d", version)
	}

	var m, t uint32
	var p uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p)
	if err != nil {
		return false, fmt.Errorf("unable to scanf m t p: %w", err)
	}

	salt, err = base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("unable to decode salt: %w", err)
	}

	hash, err = base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("unable to decode hash: %w", err)
	}

	computedHash := argon2.IDKey([]byte(password), salt, t, m, p, uint32(len(hash)))

	if subtle.ConstantTimeCompare(hash, computedHash) == 1 {
		return true, nil
	}
	return false, nil
}

func split(s, sep string) []string {
	// helper for splitting string (standard strings.Split can be used instead)
	return strings.Split(s, sep)
}
