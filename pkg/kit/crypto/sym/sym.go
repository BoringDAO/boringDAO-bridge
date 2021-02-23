package sym

import (
	"errors"
	"fmt"

	"github.com/boringdao/bridge/pkg/kit/crypto"
)

var (
	aesKeyLengthError = errors.New("the secret len must be 32")
)

// GenerateSymKey generates a new aes256 key or 3des key
func GenerateSymKey(opt crypto.KeyType, key []byte) (crypto.SymmetricKey, error) {
	switch opt {
	case crypto.AES:
		if len(key) != 32 {
			return nil, aesKeyLengthError
		}
		return &AESKey{key: key}, nil
	case crypto.ThirdDES:
		return &ThirdDESKey{key: key}, nil
	default:
		return nil, fmt.Errorf("wrong symmetric algorithm")
	}
}
