package kmgRand

import (
	"crypto/rand"
	"encoding/hex"
)

func MustCryptoRandBytes(length int) []byte {
	buf := make([]byte, length)
	_, err := rand.Read(buf)
	if err != nil {
		panic(err)
	}
	return buf
}

//读出给定长度的加密的已经Hex过的字符串(结果字符串就是那么长)
func MustCryptoRandToHex(length int) string {
	readLen := length/2 + length%2
	readbuf := make([]byte, readLen)
	_, err := rand.Read(readbuf[:readLen])
	if err != nil {
		panic(err)
	}
	outBuf := make([]byte, readLen*2)
	hex.Encode(outBuf, readbuf[:readLen])
	return string(outBuf[:length])
}

// (长度62)
const alphaNumMap = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func MustCryptoRandToAlphaNum(length int) string {
	return MustCryptoRandFromByteList(length, alphaNumMap)
}

//里面没有大小写问题,没有ilIL1问题 没有 0oO 问题. (长度31)
const realableAlphaNumMap = "23456789abcdefghjkmnpqrstuvwxyz"

func MustCryptoRandToReadableAlphaNum(length int) string {
	return MustCryptoRandFromByteList(length, realableAlphaNumMap)
}

const numMap = "0123456789"

func MustCryptoRandToNum(length int) string {
	return MustCryptoRandFromByteList(length, numMap)
}

func MustCryptoRandFromByteList(length int, list string) string {
	var bytes = make([]byte, 2*length)
	var outBytes = make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}
	mapLen := len(list)
	for i := 0; i < length; i++ {
		outBytes[i] = list[(int(bytes[2*i])*256+int(bytes[2*i+1]))%(mapLen)]
	}
	return string(outBytes)
}
