package random

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
)

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	iInt64 := i.Int64()
	if iInt64 < min {
		iInt64 = RandInt64(min, max)
	}
	return iInt64
}

func GetRandomNumString(n int) string {
	const alphanum = "012356789"
	lenAl := len(alphanum)
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, bItem := range bytes {
		tm := bItem % byte(lenAl)
		bytes[i] = alphanum[tm]
	}
	return string(bytes)
}

func GenerateRandomSalt(size int) string {
	var salt = make([]byte, size)
	_, err := rand.Read(salt[:])
	if err != nil {
		panic(err)
	}
	return base64.RawStdEncoding.EncodeToString(salt)
}
