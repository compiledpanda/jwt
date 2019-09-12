package internal

import "encoding/pem"

func dePem(key []byte) []byte {
	if len(key) > 0 && key[0] == '-' {
		block, _ := pem.Decode(key)
		if block == nil {
			return []byte{}
		}
		return block.Bytes
	}
	return key
}
