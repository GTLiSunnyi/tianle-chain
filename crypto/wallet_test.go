package crypto

import (
	"testing"
)

func TestExportAndImport(t *testing.T) {
	wallet1 := NewWallet()
	wallet1.ExportWallet("./", "123456")

	encryptData, err := wallet1.Encrypt([]byte("tianle"))
	if err != nil {
		t.Log(err)
	}

	signData, err := wallet1.Sign([]byte("gutianle"))
	if err != nil {
		t.Log(err)
	}

	wallet2 := ImportWallet("./", "123456")

	decryptData, err := wallet2.Decrypt(encryptData)
	if err != nil {
		t.Log(err)
	}
	if string(decryptData) != "tianle" {
		t.Fail()
	}

	if !wallet2.Verify([]byte("gutianle"), signData) {
		t.Fail()
	}
}
