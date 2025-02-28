package runauthorization

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"path"

	"github.com/Knubsen/vine/internal/state"
	"github.com/Knubsen/vine/internal/utils"
)

func SetAuthKey() error {
	vineSecretKey := path.Join(state.Location.Auth, ".vine_secret_key")

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}

	encodedKey := base64.StdEncoding.EncodeToString(key)

	err = os.WriteFile(vineSecretKey, []byte(encodedKey), 0644)
	if err != nil {
		return err
	}

	return nil
}

func generateAuthToken(timestamp string) (string, error) {
	vineSecretKey := path.Join(state.Location.Auth, ".vine_secret_key")

	key, err := utils.ReadLines(vineSecretKey)
	if err != nil {
		return "", err
	}
	if len(key) != 1 || key[0] == "" {
		return "", fmt.Errorf("vine secret key is not set correctly")
	}

	h := hmac.New(sha256.New, []byte(key[0]))
	h.Write([]byte(timestamp))
	return hex.EncodeToString(h.Sum(nil)), nil
}
