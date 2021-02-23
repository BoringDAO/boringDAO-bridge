package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/asn1"
	"fmt"

	"github.com/boringdao/bridge/pkg/kit/crypto"
	"github.com/boringdao/bridge/pkg/kit/types"
	"github.com/btcsuite/btcd/btcec"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
)

// PublicKey ECDSA public key.
// never new(PublicKey), use NewPublicKey()
type PublicKey struct {
	K *ecdsa.PublicKey
}

func NewPublicKey(k ecdsa.PublicKey) (*PublicKey, error) {
	switch k.Curve {
	case elliptic.P256(), elliptic.P384(), elliptic.P521(), btcec.S256():
		break
	default:
		return nil, fmt.Errorf("unsupported ecdsa curve option")
	}
	return &PublicKey{K: &k}, nil
}

func UnmarshalPublicKey(data []byte, opt crypto.KeyType) (crypto.PublicKey, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty public key data")
	}
	var (
		pub *ecdsa.PublicKey
		ok  bool
	)
	if opt == crypto.Secp256k1 {
		rawPub, err := btcec.ParsePubKey(data, btcec.S256())
		if err != nil {
			return nil, err
		}

		pub = rawPub.ToECDSA()
	} else {
		pubInfo, err := x509.ParsePKIXPublicKey(data)
		if err != nil {
			return nil, err
		}

		pub, ok = pubInfo.(*ecdsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("not ecdsa public key")
		}
	}

	key := &PublicKey{K: pub}
	switch opt {
	case crypto.ECDSA_P256, crypto.ECDSA_P384, crypto.ECDSA_P521, crypto.Secp256k1:
		key.K.Curve = pub.Curve
	default:
		return nil, fmt.Errorf("not supported curve")
	}
	return key, nil
}

func Unmarshal(data []byte) (crypto.PrivateKey, error) {
	priv, err := x509.ParseECPrivateKey(data)
	if err != nil {
		return nil, err
	}

	return &PrivateKey{curve: crypto.ECDSA_P256, K: priv}, nil
}

func (pub *PublicKey) Bytes() ([]byte, error) {
	if pub.K == nil {
		return nil, fmt.Errorf("ECDSAPublicKey.K is nil, please invoke FromBytes()")
	}

	if pub.Type() == crypto.Secp256k1 {
		rawKey := (*btcec.PublicKey)(pub.K).SerializeCompressed()
		return rawKey, nil
	}

	return x509.MarshalPKIXPublicKey(pub.K)
}

func (pub *PublicKey) Address() (types.Address, error) {
	data := elliptic.Marshal(pub.K.Curve, pub.K.X, pub.K.Y)

	ret := ecrypto.Keccak256(data[1:])

	return types.Bytes2Address(ret[12:]), nil
}

func (pub *PublicKey) Verify(digest []byte, sig []byte) (bool, error) {
	if sig == nil {
		return false, fmt.Errorf("nil signature")
	}

	sigStruct := &Sig{}
	if _, err := asn1.Unmarshal(sig, sigStruct); err != nil {
		return false, err
	}

	if !ecdsa.Verify(pub.K, digest, sigStruct.R, sigStruct.S) {
		return false, fmt.Errorf("invalid signature")
	}

	return true, nil
}

func (pub *PublicKey) Type() crypto.KeyType {
	switch pub.K.Curve {
	case elliptic.P256():
		return crypto.ECDSA_P256
	case elliptic.P384():
		return crypto.ECDSA_P384
	case elliptic.P521():
		return crypto.ECDSA_P521
	case btcec.S256():
		return crypto.Secp256k1
	}
	return -1
}
