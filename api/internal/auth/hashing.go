package auth

import (
	"bytes"
	"crypto/rand"
	"golang.org/x/crypto/argon2"
)

type IHashingService interface {
	HashPassword(password string) []byte
	CheckPassword(password string, storedHashAndSalt []byte) bool
}

type HashingService struct {
	mem         uint32
	time        uint32
	parallelism uint8
	hashLen     uint32
	saltLen     uint32
}

func NewHashingService() *HashingService {
	return &HashingService{
		// params according to https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html#argon2id
		mem:         12288,
		time:        3,
		parallelism: 1,
		hashLen:     64,
		saltLen:     32,
	}
}
func (hs *HashingService) HashPassword(password string) []byte {

	salt := hs.generateSalt()

	hash := hs.hashPassword(password, salt)
	return hash
}
func (hs *HashingService) hashPassword(password string, salt []byte) []byte {
	hash := argon2.IDKey([]byte(password), salt, hs.time, hs.mem, hs.parallelism, hs.hashLen)

	hashAndSalt := append(hash, salt...)
	return hashAndSalt
}

func (hs *HashingService) CheckPassword(password string, storedHashAndSalt []byte) bool {
	storedHash, storedSalt := storedHashAndSalt[0:hs.hashLen], storedHashAndSalt[hs.hashLen:]

	// call the hash fn, but strip the salt off.
	hash := hs.hashPassword(password, storedSalt)[0:hs.hashLen]

	return bytes.Equal(storedHash, hash)
}

func (hs *HashingService) generateSalt() []byte {
	salt := make([]byte, hs.saltLen)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	return salt

}
