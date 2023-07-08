package encode

func BoolToBytes(v bool) []byte {
	if v {
		return []byte{0x01}
	} else {
		return []byte{0x00}
	}
}
