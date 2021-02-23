package asym

import (
	"crypto/sha256"
	"os"
	"path/filepath"
	"testing"

	"github.com/boringdao/bridge/pkg/kit/crypto"
	"github.com/boringdao/bridge/pkg/kit/crypto/asym/ecdsa"
	"github.com/stretchr/testify/require"
)

func TestSignAndVerifyPass(t *testing.T) {
	testSignAndVerify(t, crypto.ECDSA_P256)
	testSignAndVerify(t, crypto.ECDSA_P384)
	testSignAndVerify(t, crypto.ECDSA_P521)
	testSignAndVerify(t, crypto.Secp256k1)
}

func TestSignAndFail(t *testing.T) {
	testSignAndVerifyFail(t, crypto.ECDSA_P256)
	testSignAndVerifyFail(t, crypto.ECDSA_P384)
	testSignAndVerifyFail(t, crypto.ECDSA_P521)
	testSignAndVerifyFail(t, crypto.Secp256k1)
}

func TestStorePrivateKey(t *testing.T) {
	testStore(t, crypto.ECDSA_P256)
	testStore(t, crypto.ECDSA_P384)
	testStore(t, crypto.ECDSA_P521)
	testStore(t, crypto.Secp256k1)
}

func testStore(t *testing.T, opt crypto.KeyType) {
	key, err := ecdsa.New(opt)
	require.Nil(t, err)

	keyFile := filepath.Join(os.TempDir(), "priv.json")

	err = StorePrivateKey(key, keyFile, "key")
	require.Nil(t, err)

	newKey, err := RestorePrivateKey(keyFile, "key")
	require.Nil(t, err)
	require.NotNil(t, newKey)

	address, err := newKey.PublicKey().Address()
	require.Nil(t, err)

	oldAddr, err := key.PublicKey().Address()
	require.Nil(t, err)

	require.EqualValues(t, oldAddr, address)
}

func testSignAndVerify(t *testing.T, opt crypto.KeyType) {
	digest := sha256.Sum256([]byte("hyperchain"))

	priv, err := GenerateKeyPair(opt)
	require.Nil(t, err)

	addr, err := priv.PublicKey().Address()
	require.Nil(t, err)

	sig, err := priv.Sign(digest[:])
	require.Nil(t, err)

	b, err := Verify(opt, sig, digest[:], addr)
	require.Nil(t, err)
	require.Equal(t, true, b)
}

func testSignAndVerifyFail(t *testing.T, opt crypto.KeyType) {
	digest := sha256.Sum256([]byte("hyperchain"))

	priv, err := GenerateKeyPair(opt)
	require.Nil(t, err)

	addr, err := priv.PublicKey().Address()
	require.Nil(t, err)

	sig, err := priv.Sign(digest[:])
	require.Nil(t, err)

	wrongDigest := sha256.Sum256([]byte("hypercha1n"))

	b, err := Verify(opt, sig, wrongDigest[:], addr)
	require.NotNil(t, err)
	require.Equal(t, false, b)
}
