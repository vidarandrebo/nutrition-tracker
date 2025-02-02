package auth

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

type AuthService struct {
	mem         uint32
	time        uint32
	parallelism uint8
	hashLen     uint32
	saltLen     uint32
}

func NewAuthService() *AuthService {
	return &AuthService{
		// params according to https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
		mem:         12288,
		time:        3,
		parallelism: 1,
		hashLen:     64,
		saltLen:     32,
	}
}
func (as *AuthService) HashPassword(password string, salt []byte) []byte {
	hash := argon2.IDKey([]byte(password), salt, as.time, as.mem, as.parallelism, as.hashLen)

	hashAndSalt := append(hash, salt...)
	return hashAndSalt
}

func (as *AuthService) CheckPassword(password string, storedHashAndSalt []byte) bool {
	storedHash, storedSalt := storedHashAndSalt[0:as.hashLen], storedHashAndSalt[as.hashLen:]

	// call the hash fn, but strip the salt off.
	hash := as.HashPassword(password, storedSalt)[0:as.hashLen]

	return bytes.Equal(storedHash, hash)
}

func (as *AuthService) GenerateSalt() []byte {
	salt := make([]byte, as.saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt

}
