package asym

import (
	crypto2 "crypto"
	ecdsa2 "crypto/ecdsa"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/asn1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/boringdao/bridge/pkg/kit/crypto"
	"github.com/boringdao/bridge/pkg/kit/crypto/asym/ecdsa"
	"github.com/boringdao/bridge/pkg/kit/crypto/sym"
	"github.com/boringdao/bridge/pkg/kit/types"
)

func GenerateKeyPair(opt crypto.KeyType) (crypto.PrivateKey, error) {
	switch opt {
	case crypto.RSA:
		return nil, fmt.Errorf("don`t support rsa algorithm currently")
	case crypto.ECDSA_P256, crypto.ECDSA_P384, crypto.ECDSA_P521, crypto.Secp256k1:
		return ecdsa.New(opt)
	case crypto.Ed25519:
		return nil, fmt.Errorf("don`t support ed25519 algorithm currently")
	default:
		return nil, fmt.Errorf("wront algorithm type")
	}
}

func Verify(opt crypto.KeyType, sig, digest []byte, from types.Address) (bool, error) {
	switch opt {
	case crypto.RSA:
		return false, fmt.Errorf("don`t support rsa algorithm currently")
	case crypto.ECDSA_P256, crypto.ECDSA_P384, crypto.ECDSA_P521, crypto.Secp256k1:
		sigStuct := &ecdsa.Sig{}
		_, err := asn1.Unmarshal(sig, sigStuct)
		if err != nil {
			return false, err
		}

		pubkey, err := ecdsa.UnmarshalPublicKey(sigStuct.Pub, opt)
		if err != nil {
			return false, err
		}

		expected, err := pubkey.Address()
		if err != nil {
			return false, err
		}

		if expected.Hex() != from.Hex() {
			return false, fmt.Errorf("wrong singer for this signature")
		}
		return pubkey.Verify(digest, sig)
	case crypto.Ed25519:
		return false, fmt.Errorf("don`t support ed25519 algorithm currently")
	default:
		return false, fmt.Errorf("wront algorithm type")
	}
}

// PrivateKeyFromStdKey convert golang standard crypto key to our private key
func PrivateKeyFromStdKey(priv crypto2.PrivateKey) (crypto.PrivateKey, error) {
	switch key := priv.(type) {
	case *ecdsa2.PrivateKey:
		return ecdsa.NewWithCryptoKey(key)
	case *ed25519.PrivateKey:
		return nil, fmt.Errorf("don't support this algorithm")
	default:
		return nil, fmt.Errorf("don't support this algorithm")
	}
}

func PubKeyFromStdKey(pub crypto2.PublicKey) (crypto.PublicKey, error) {
	switch key := pub.(type) {
	case *ecdsa2.PublicKey:
		return ecdsa.NewPublicKey(*key)
	case *ed25519.PublicKey:
		return nil, fmt.Errorf("don't support this algorithm")
	default:
		return nil, fmt.Errorf("don't support this algorithm")
	}
}

// PrivKeyToStdKey convert our crypto private key to golang standard ecdsa private key
func PrivKeyToStdKey(priv crypto.PrivateKey) (ecdsa2.PrivateKey, error) {
	switch p := priv.(type) {
	case *ecdsa.PrivateKey:
		return *p.K, nil
	default:
		return ecdsa2.PrivateKey{}, fmt.Errorf("don't support this algorithm")
	}
}

func PubKeyToStdKey(pub crypto.PublicKey) (crypto2.PublicKey, error) {
	switch key := pub.(type) {
	case *ecdsa.PublicKey:
		return key.K, nil
	default:
		return ecdsa2.PublicKey{}, fmt.Errorf("don't support this algorithm")
	}
}

func StorePrivateKey(priv crypto.PrivateKey, keyFilePath, password string) error {
	keyStore, err := GenKeyStore(priv, password)
	if err != nil {
		return err
	}

	filePtr, err := os.Create(keyFilePath)
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(keyStore, "", " ")
	if err != nil {
		return err
	}

	_, err = filePtr.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func GenKeyStore(priv crypto.PrivateKey, password string) (*crypto.KeyStore, error) {
	var cipherText string

	privBytes, err := priv.Bytes()
	if err != nil {
		return nil, err
	}
	if password != "" {
		hash := sha256.Sum256([]byte(password))
		aesKey, err := sym.GenerateSymKey(crypto.AES, hash[:])
		if err != nil {
			return nil, err
		}

		encrypted, err := aesKey.Encrypt(privBytes)
		if err != nil {
			return nil, err
		}

		cipherText = hex.EncodeToString(encrypted)
	} else {
		cipherText = hex.EncodeToString(privBytes)
	}

	return &crypto.KeyStore{
		Type: priv.Type(),
		Cipher: &crypto.CipherKey{
			Cipher: "AES-256",
			Data:   cipherText,
		},
	}, nil
}

func RestorePrivateKey(keyFilePath, password string) (crypto.PrivateKey, error) {
	data, err := ioutil.ReadFile(keyFilePath)
	if err != nil {
		return nil, err
	}
	keyStore := &crypto.KeyStore{}
	if err := json.Unmarshal(data, keyStore); err != nil {
		return nil, err
	}

	rawBytes, err := hex.DecodeString(keyStore.Cipher.Data)
	if err != nil {
		return nil, err
	}

	if password != "" {
		hash := sha256.Sum256([]byte(password))
		aesKey, err := sym.GenerateSymKey(crypto.AES, hash[:])
		if err != nil {
			return nil, err
		}

		rawBytes, err = aesKey.Decrypt(rawBytes)
		if err != nil {
			return nil, err
		}
	}

	switch keyStore.Type {
	case crypto.ECDSA_P256, crypto.ECDSA_P384, crypto.ECDSA_P521, crypto.Secp256k1:
		return ecdsa.UnmarshalPrivateKey(rawBytes, keyStore.Type)
	case crypto.Ed25519, crypto.RSA:
		return nil, fmt.Errorf("don't support this private key")
	default:
		return nil, fmt.Errorf("don't support this private key")
	}
}
