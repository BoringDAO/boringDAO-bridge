package ecdh

import "github.com/boringdao/bridge/pkg/kit/crypto"

type KeyExchange interface {
	// Check returns a non-nil error if the peers public key cannot used for the
	// key exchange - for instance the public key isn't a point on the elliptic curve.
	// It's recommended to check peer's public key before computing the secret.
	Check(peerPubkey []byte) error

	// ComputeSecret returns the secret value computed from the given private key
	// and the peers public key.
	ComputeSecret(privkey crypto.PrivateKey, peerPubkey []byte) (secret []byte, err error)
}
