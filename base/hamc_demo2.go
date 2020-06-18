package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// 将字符串加盐后，生成sha256加密字符串，
func GenerateHmacString(str, salt string) (hmacStr string) {
	hash := hmac.New(sha256.New, []byte(salt))
	hash.Write([]byte(str))
	hmacStr = hex.EncodeToString(hash.Sum(nil))
	return
}

// 比较加密字符串是否相等
func CheckHmac(str, hmacStr, salt string) bool {
	mac := hmac.New(sha256.New, []byte(salt))
	mac.Write([]byte(str))
	expectedMAC := mac.Sum(nil)
	v, err := hex.DecodeString(hmacStr)
	if err != nil {
		return false
	}
	return hmac.Equal(expectedMAC, v)
}

func main() {
	hmacS := GenerateHmacString("ymdd", "JabH4hoCzVY8tCvY")
	fmt.Println(hmacS)
	//
	//b := CheckHmac("ymdd", "e1c336f4db2909ff93af9d840d01cd617929af02a98596403d8a1ec4b524e8ee", "JabH4hoCzVY8tCvY")
	//fmt.Println(b)

}
