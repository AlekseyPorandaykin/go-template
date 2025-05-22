package secrets

import (
	"github.com/dromara/dongle"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"strconv"
)

func HiddenPassword(pass string) string {
	hiddenStr := "**********"
	var result []byte
	passData := []byte(pass)
	if len(passData) < 5 {
		return hiddenStr
	}
	if len(passData) < 10 {
		result = append(result, passData[:2]...)
		result = append(result, []byte(hiddenStr)...)
		result = append(result, passData[len(passData)-2:]...)
		return string(result)
	}
	result = append(result, passData[:4]...)
	result = append(result, []byte(hiddenStr)...)
	result = append(result, passData[len(passData)-3:]...)
	return string(result)
}

func HiddenUid(uid uuid.UUID) string {
	var res int
	for _, b := range uid {
		res += int(b)
	}
	return strconv.Itoa(res)
}

func DecodePassword(data string) string {
	return dongle.Decrypt.FromRawString(data).ByRsa([]byte(viper.GetString("secret.rsa.private_key"))).ToString()
}

func EncodePassword(data string) string {
	return dongle.Encrypt.FromString(data).ByRsa([]byte(viper.GetString("secret.rsa.public_key"))).ToRawString()
}
