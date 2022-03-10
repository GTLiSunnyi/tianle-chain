package crypto

import "testing"

func TestSha256(t *testing.T) {
	hash := Sha256([]byte("tianle"))
	if hash != "609a686b33d32a8f877bdfb3762e4c67680d536f3eae99ea2130500ea9228c24" {
		t.Fail()
	}
}
