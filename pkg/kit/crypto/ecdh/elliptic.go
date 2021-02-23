package ecdh

import (
	"crypto/elliptic"
	"fmt"
	"math/big"

	"github.com/boringdao/bridge/pkg/kit/crypto"
	"github.com/boringdao/bridge/pkg/kit/crypto/asym/ecdsa"
)

type ellipticECDH struct {
	curve elliptic.Curve
}

func NewEllipticECDH(c elliptic.Curve) (KeyExchange, error) {
	if c == nil {
		return nil, fmt.Errorf("invalid curve")
	}

	return ellipticECDH{curve: c}, nil
}

func (e ellipticECDH) Check(peerPubkey []byte) error {
	if len(peerPubkey) == 0 {
		return fmt.Errorf("empty public key byte")
	}
	x, y, err := getXYFromPub(peerPubkey)
	if err != nil {
		return err
	}
	if !e.curve.IsOnCurve(x, y) {
		return fmt.Errorf("peer's public key is not on curve")
	}

	return nil
}

func (e ellipticECDH) ComputeSecret(privkey crypto.PrivateKey, peerPubkey []byte) ([]byte, error) {
	err := e.Check(peerPubkey)
	if err != nil {
		return nil, err
	}

	x, y, err := getXYFromPub(peerPubkey)
	if err != nil {
		return nil, err
	}

	privBytes, err := privkey.Bytes()
	if err != nil {
		return nil, err
	}

	sX, _ := e.curve.ScalarMult(x, y, privBytes)
	secret := sX.Bytes()

	return secret, nil
}

func getXYFromPub(pub []byte) (*big.Int, *big.Int, error) {
	key, err := ecdsa.UnmarshalPublicKey(pub, crypto.Secp256k1)
	if err != nil {
		return nil, nil, err
	}

	pubKey := key.(*ecdsa.PublicKey)
	return pubKey.K.X, pubKey.K.Y, nil
}
