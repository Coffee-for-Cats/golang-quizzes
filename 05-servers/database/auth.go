package db

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func hash_credentials(pass, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(pass))
	hash.Write([]byte(salt))

	return hex.EncodeToString(hash.Sum(nil))
}

func RegisterUser(name string, pass string) (err error) {

	salt_bytes := make([]byte, 20)
	_, err1 := rand.Read(salt_bytes)
	if err1 != nil {
		return err1
	}

	salt := hex.EncodeToString(salt_bytes)

	hash_str := hash_credentials(pass, salt)

	_, err = Use().Exec(`INSERT INTO users
		(name, salt, hash)
		VALUES ($1, $2, $3)`,
		name, salt, hash_str,
	)

	if err != nil {
		return err
	}

	return nil
}

func VerifyUser(name string, pass string) bool {

	row := Use().QueryRow(`SELECT salt, hash
		FROM users
		WHERE name = $1
	`, name)

	var salt, hash string
	err := row.Scan(&salt, &hash)
	if err != nil {
		return false // user not found
	}

	expected_hash := hash_credentials(pass, salt)

	if expected_hash != hash {
		return false // incorrect password
	}

	return true
}
