package db

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func RegisterUser(name string, pass string) (err error) {

	salt_bytes := make([]byte, 20)
	_, err1 := rand.Read(salt_bytes)
	if err1 != nil {
		return err1
	}

	salt := base64.URLEncoding.EncodeToString(salt_bytes)

	hash := sha256.New()
	hash.Write([]byte(pass))
	hash.Write([]byte(salt))

	hash_str := string(hash.Sum(nil))

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
		WHERE name == $1
	`, name)

	var salt, hash string
	err := row.Scan(&salt, &hash)
	if err != nil {
		return false // user not found
	}

	expected_hash := sha256.New()
	expected_hash.Write([]byte(pass))
	expected_hash.Write([]byte(salt))

	expected_hash_str := string(expected_hash.Sum(nil))
	if expected_hash_str != hash {
		return false // incorrect password
	}

	return true
}
