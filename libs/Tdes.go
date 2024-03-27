package libs

import (
	"encoding/hex"
	"fmt"
	"mqtt/help"

	"github.com/forgoer/openssl"
)

func TdesEncrypt(key, data []byte) (byteData []byte, err error) {
	iv := make([]byte, 8)
	if byteData, err = openssl.Des3CBCEncrypt(data, key, iv, openssl.PKCS7_PADDING); err != nil {
		help.ErrorLog(fmt.Sprintf("TdesEncrypt error:%s\n, key:%s, data:%s", err.Error(), hex.EncodeToString(key), hex.EncodeToString(data)))
	}
	return
}

func TdesDescrypt(key, data []byte) (byteData []byte, err error) {
	iv := make([]byte, 8)
	if byteData, err = openssl.Des3CBCDecrypt(data, key, iv, openssl.PKCS7_PADDING); err != nil {
		help.ErrorLog(fmt.Sprintf("TdesEncrypt error:%s\n, key:%s, data:%s", err.Error(), hex.EncodeToString(key), hex.EncodeToString(data)))
	}
	return
}
