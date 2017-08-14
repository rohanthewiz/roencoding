package roencoding

import (
	"github.com/pierrec/xxHash/xxHash64"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"crypto/sha1"
	"time"
)

func XXHash(str string) string {
	const random_int = 49213745817371
	return fmt.Sprintf("%x",xxHash64.Checksum([]byte(str), random_int))
}

func HashMD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Sha1Random() string {
	return fmt.Sprintf("%x", sha1.Sum([]byte("%$" + time.Now().String() + "e{")))
}
