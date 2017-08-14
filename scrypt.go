package roencoding

import (
	"fmt"
	"time"
	"golang.org/x/crypto/scrypt"
	"github.com/rohanthewiz/serr"
)

func PasswordHash(password, salt string) (string, error) {
	return cryptit(8, 32, password, salt)
}

func GenSalt(in string) (string, error) {
	str := fmt.Sprintf("imrandom%szoq;lnesh)^%di!!!2#@", in, time.Now().Unix())
	return cryptit(12, 24, str)
}

func cryptit(r int, key_len int, password string, salts ...string) (string, error) {
	salt := ""
	if salts == nil {
		salt = fmt.Sprintf("00--%x", time.Now().Unix())
	} else {
		salt = salts[0]
	}
	dk, err := scrypt.Key([]byte(password), []byte(salt), 16384, r, 1, key_len)
	if err != nil {
		return "", serr.Wrap(err, "Hash function failed")
	}
	return fmt.Sprintf("%x", dk), nil
}
