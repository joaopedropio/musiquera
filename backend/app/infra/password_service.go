package infra

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"crypto/subtle"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/argon2"
)

const encodedHashFormat = "$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s"

type PasswordService interface {
	JWTAuth() *jwtauth.JWTAuth
	HashPassword(password string) (encodedHash string, err error)

	VerifyPassword(password, encodedHash string) (bool, error)
}

type passwordService struct {
	defaultParams *argon2Params
	jwtAuth       *jwtauth.JWTAuth
}

func NewPasswordService(jwtSecret string) PasswordService {
	return &passwordService{
		jwtAuth: createJwtAuth(jwtSecret),
		defaultParams: &argon2Params{
			Memory:      64 * 1024, // 64 MB
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		},
	}
}

type argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

func (s *passwordService) JWTAuth() *jwtauth.JWTAuth {
	return s.jwtAuth
}

func createJwtAuth(jwtSecret string) *jwtauth.JWTAuth {
	algorithm := "HS256"
	signKey := []byte(jwtSecret)
	return jwtauth.New(algorithm, signKey, nil)
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	return b, err
}

func (s *passwordService) HashPassword(password string) (encodedHash string, err error) {
	return hashPassword(password, s.defaultParams)
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

func (s *passwordService) VerifyPassword(password, encodedHash string) (bool, error) {
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
