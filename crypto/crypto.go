package crypto

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm2"
)

// SHA256
func Sha256(data []byte) string {
	bytes := sha256.Sum256(data)
	return hex.EncodeToString(bytes[:])
}

func GetAddressByPubKey(pubKey *sm2.PublicKey) string {
	hash := Sha256(append(pubKey.X.Bytes(), pubKey.Y.Bytes()...))
	return "0x" + Sha256([]byte(hash)[43:])
}
