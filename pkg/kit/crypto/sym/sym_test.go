package sym

import (
	"crypto/rand"
	"testing"

	"github.com/boringdao/bridge/pkg/kit/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const msg = `Qulian Technology is an international leading blockchain team with all core team members graduated from Zhejiang University, Tsinghua University and other first-class universities at home and abroad, and Academician Chen Chun of the Chinese Academy of Engineering acted as chairman of the board. The company has a team of nearly 200 people, 90% of whom are technicians, more than 10 have doctoral degrees and 140 have master's degrees. The core competitiveness of the company is Hyperchain bottom technology platform. This platform ranks first in the technical evaluation of several large and medium-sized financial institutions. It is also the first batch of bottom platforms to pass the Blockchain Standard Test of the China Electronics Standardization Institute (CESI) and China Academy of Information and Communications Technology (CAICT) of Ministry of Industry and Information Technology (MIIT). It has applied for 28 patents in blockchain related fields.`

func TestAES(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	require.Nil(t, err)
	aes, err := GenerateSymKey(crypto.AES, key)
	require.Nil(t, err)

	c, err := aes.Encrypt([]byte(msg))
	require.Nil(t, err)

	o, err := aes.Decrypt(c)
	require.Nil(t, err)

	require.Equal(t, string(o), msg)
}

func TestAESKey_Error(t *testing.T) {
	key := make([]byte, 12)
	_, err := rand.Read(key)
	require.Nil(t, err)
	_, err = GenerateSymKey(crypto.AES, key)
	require.Equal(t, aesKeyLengthError, err)
}

func TestThirdDESKey(t *testing.T) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	assert.Nil(t, err)
	des, err := GenerateSymKey(crypto.ThirdDES, key)
	assert.Nil(t, err)

	c, err := des.Encrypt([]byte(msg))
	require.Nil(t, err)

	o, err := des.Decrypt(c)
	require.Nil(t, err)

	require.Equal(t, string(o), msg)
}
