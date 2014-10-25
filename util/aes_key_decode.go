// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

import (
	"encoding/base64"
	"errors"
)

// encodedAESKey 由 a-z,A-Z,0-9 组成
func AESKeyDecode(encodedAESKey string) (AESKey []byte, err error) {
	if len(encodedAESKey) != 43 {
		err = errors.New("the length of encodedAESKey must be equal to 43")
		return
	}
	return base64.StdEncoding.DecodeString(encodedAESKey + "=")
}