package crypto

import (
	"crypto/rand"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

type Wallet struct {
	KeyPair *KeyPair
	Address string
}

func NewWallet() *Wallet {
	keyPair := NewKeyPair()

	addr := GetAddressByPubKey(keyPair.PublicKey)

	return &Wallet{&keyPair, addr}
}

// 签名
func (w *Wallet) Sign(msg []byte) ([]byte, error) {
	return w.KeyPair.PrivateKey.Sign(rand.Reader, msg, nil)
}

// 验签
func (w *Wallet) Verify(msg, data []byte) bool {
	return w.KeyPair.PublicKey.Verify(msg, data)
}

// 加密
func (w *Wallet) Encrypt(msg []byte) ([]byte, error) {
	return w.KeyPair.PublicKey.EncryptAsn1(msg, rand.Reader)
}

// 解密
func (w *Wallet) Decrypt(data []byte) ([]byte, error) {
	return w.KeyPair.PrivateKey.DecryptAsn1(data)
}

// 导出 wallet
func (w *Wallet) ExportWallet(path, pwd string) {
	w.ExportPrivKey(path, pwd)
	w.ExportPubKey(path)
	w.ExportAddress(path)
}

// 导出私钥
func (w *Wallet) ExportPrivKey(path, pwd string) {
	data, err := x509.WritePrivateKeyToPem(w.KeyPair.PrivateKey, []byte(pwd))
	if err != nil {
		log.Fatal(err)
	}

	block := &pem.Block{
		Type:  "SM2 PRIVATE KEY",
		Bytes: data,
	}

	file, err := os.Create(path + "/private.pem")
	if err != nil {
		log.Fatal(err)
	}

	err = pem.Encode(file, block)
	if err != nil {
		log.Fatal(err)
	}
}

// 导出公钥
func (w *Wallet) ExportPubKey(path string) {
	data, err := x509.WritePublicKeyToPem(w.KeyPair.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	block := &pem.Block{
		Type:  "SM2 PUBLIC KEY",
		Bytes: data,
	}

	file, err := os.Create(path + "/public.pem")
	if err != nil {
		log.Fatal(err)
	}

	err = pem.Encode(file, block)
	if err != nil {
		log.Fatal(err)
	}
}

// 导出地址
func (w *Wallet) ExportAddress(path string) {
	file, err := os.Create(path + "/address")
	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write([]byte(w.Address))
	if err != nil {
		log.Fatal(err)
	}
}

// 导入 wallet
func ImportWallet(path, pwd string) *Wallet {
	privKey := ImportPrivKey(path, pwd)
	pubKey := ImportPubKey(path)
	address := ImportAddress(path)

	return &Wallet{&KeyPair{privKey, pubKey}, address}
}

// 导入私钥
func ImportPrivKey(path, pwd string) *sm2.PrivateKey {
	data, err := ioutil.ReadFile(path + "/private.pem")
	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		log.Fatal("pem.Decode error")
	}

	privKey, err := x509.ReadPrivateKeyFromPem(block.Bytes, []byte(pwd))
	if err != nil {
		log.Fatal(err)
	}

	return privKey
}

// 导入公钥
func ImportPubKey(path string) *sm2.PublicKey {
	data, err := ioutil.ReadFile(path + "/public.pem")
	if err != nil {
		log.Fatal(err)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		log.Fatal("pem.Decode error")
	}

	pubKey, err := x509.ReadPublicKeyFromPem(block.Bytes)
	if err != nil {
		log.Fatal(err)
	}

	return pubKey
}

// 导入地址
func ImportAddress(path string) string {
	address, err := ioutil.ReadFile(path + "/address")
	if err != nil {
		log.Fatal(err)
	}

	return string(address)
}
