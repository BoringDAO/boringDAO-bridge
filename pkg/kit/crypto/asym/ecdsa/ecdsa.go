package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/asn1"
	"fmt"
	"math/big"

	"github.com/boringdao/bridge/pkg/kit/crypto"
	"github.com/btcsuite/btcd/btcec"
)

var _ crypto.PrivateKey = (*PrivateKey)(nil)
var _ crypto.PublicKey = (*PublicKey)(nil)

// PrivateKey ECDSA private key.
// never new(PrivateKey), use NewPrivateKey()
type PrivateKey struct {
	curve crypto.KeyType
	K     *ecdsa.PrivateKey
}

// ECDSASig holds the r and s values of an ECDSA signature
type Sig struct {
	Pub []byte   `json:"pub"`
	R   *big.Int `json:"r"`
	S   *big.Int `json:"s"`
}

// New generate a ecdsa private key,input is algorithm type
func New(opt crypto.KeyType) (crypto.PrivateKey, error) {
	switch opt {
	case crypto.Secp256k1:
		pri, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri, curve: crypto.Secp256k1}, nil
	case crypto.ECDSA_P256:
		pri, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri, curve: crypto.ECDSA_P256}, nil
	case crypto.ECDSA_P384:
		pri, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri, curve: crypto.ECDSA_P384}, nil
	case crypto.ECDSA_P521:
		pri, err := ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
		if err != nil {
			return nil, err
		}

		return &PrivateKey{K: pri, curve: crypto.ECDSA_P521}, nil
	}
	return nil, fmt.Errorf("wrong curve option")
}

func NewWithCryptoKey(priv *ecdsa.PrivateKey) (crypto.PrivateKey, error) {
	newPrivKey := &PrivateKey{}
	switch priv.Curve {
	case elliptic.P256():
		newPrivKey.curve = crypto.ECDSA_P256
	case elliptic.P384():
		newPrivKey.curve = crypto.ECDSA_P384
	case elliptic.P521():
		newPrivKey.curve = crypto.ECDSA_P521
	case btcec.S256():
		newPrivKey.curve = crypto.Secp256k1
	}

	newPrivKey.K = priv
	return newPrivKey, nil
}

// Bytes returns a serialized, storable representation of this key
func (priv *PrivateKey) Bytes() ([]byte, error) {
	if priv.K == nil {
		return nil, fmt.Errorf("ECDSAPrivateKey.K is nil, please invoke FromBytes()")
	}

	if priv.Type() == crypto.Secp256k1 {
		rawKey := (*btcec.PrivateKey)(priv.K).Serialize()
		return rawKey, nil
	}
	return x509.MarshalECPrivateKey(priv.K)
}

func (priv *PrivateKey) PublicKey() crypto.PublicKey {
	return &PublicKey{K: &priv.K.PublicKey}
}

func (priv *PrivateKey) Sign(digest []byte) ([]byte, error) {
	r, s, err := ecdsa.Sign(rand.Reader, priv.K, digest[:])
	if err != nil {
		return nil, err
	}

	pubBytes, err := priv.PublicKey().Bytes()
	if err != nil {
		return nil, err
	}

	return asn1.Marshal(Sig{Pub: pubBytes, R: r, S: s})
}

func (priv *PrivateKey) Type() crypto.KeyType {
	return priv.curve
}

func UnmarshalPrivateKey(data []byte, opt crypto.KeyType) (*PrivateKey, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty private key data")
	}

	var (
		priv *ecdsa.PrivateKey
		err  error
	)
	if opt == crypto.Secp256k1 {
		Secp256k1Key, _ := btcec.PrivKeyFromBytes(btcec.S256(), data)
		priv = Secp256k1Key.ToECDSA()
	} else {
		priv, err = x509.ParseECPrivateKey(data)
		if err != nil {
			return nil, err
		}
	}

	pri := &PrivateKey{K: priv, curve: opt}
	switch opt {
	case crypto.ECDSA_P256, crypto.ECDSA_P384, crypto.ECDSA_P521, crypto.Secp256k1:
		pri.K.Curve = priv.Curve
	default:
		return nil, fmt.Errorf("not supported curve")
	}

	return pri, nil
}
