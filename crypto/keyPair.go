package crypto

import (
	"crypto/rand"
	"log"

	"github.com/tjfoc/gmsm/sm2"
)

type KeyPair struct {
	PrivateKey *sm2.PrivateKey
	PublicKey  *sm2.PublicKey
}

// 生成密钥对
func NewKeyPair() KeyPair {
	privKey, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	pubKey := &privKey.PublicKey

	return KeyPair{privKey, pubKey}
}
