package _crypto

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password string, cost ...int) (r string, err error) {
	c := bcrypt.DefaultCost
	l := len(cost)
	if l > 0 {
		if l > 1 {
			err = fmt.Errorf("cost only accept 1 int value")
			return
		}
		if cost[0] >= bcrypt.MinCost {
			c = cost[0]
		}
	}
	_r, err := bcrypt.GenerateFromPassword([]byte(password), c)
	if err != nil {
		return
	}
	r = string(_r)
	return
}

func BcryptVerify(hashPassword, password string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
