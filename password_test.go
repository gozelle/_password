package _crypto

import "testing"

func TestHashPassword(t *testing.T) {
	pass := "123456"
	r, err := BcryptHash("123456")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("hashed password:", r)
	err = BcryptVerify(r, pass)
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log("密码校验通过")
	}
	err = BcryptVerify(r, "12345")
	if err != nil {
		t.Error(err)
		return
	} else {
		t.Log("密码校验通过")
	}
}
