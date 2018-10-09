package base62

import "bytes"

type Base62 struct {
	sourceBase int
	targetBase int
	Message    string
	alphabet   []byte
	lookup     []byte
}

const BASE  = "06TLetOFlcijmbuDH4NRI5Vs3MZXfxnYB1SgQpy8qPAordvJhaG7EkUC92wWKz"

func (base62 *Base62) Encode(message string) (encoded string) {
	messageBytes := TranslateStringToByteArray(message)
	encoded = TranslateByteArrayToString(translate(Convert(messageBytes, base62.sourceBase, base62.targetBase),base62.alphabet))
	return
}

func (base62 *Base62) Decode(encoded string) (message string) {
	encodedBytes := translate(TranslateStringToByteArray(encoded),base62.lookup)
	message = TranslateByteArrayToString(Convert(encodedBytes,base62.targetBase,base62.sourceBase))
	return
}

func (base62 *Base62) createLookupTable() {
	base62.lookup = make([]byte,256)
	for k, v := range base62.alphabet {
		base62.lookup[v] = byte(k)
	}
}

func GetBase62Instance() *Base62 {
	base62 := &Base62{
		sourceBase: 256,
		targetBase: 62,
		alphabet:   TranslateStringToByteArray(BASE),
	}
	base62.createLookupTable()
	return base62
}

func TranslateStringToByteArray(str string) (alps []byte) {
	byteBuffer := bytes.Buffer{}
	byteBuffer.WriteString(str)
	return byteBuffer.Bytes()
}

func TranslateByteArrayToString(datas []byte) (str string) {
	return string(datas)
}

/**

 */
func Convert(data []byte, sourceBase int, targetBase int) (ret []byte) {
	source := make([]int, len(data))
	for k, v := range data {
		source[k] = int(v)
	}
	var remainder int
	for len(source) > 0 {
		var quotient = make([]int, 0, 10)
		remainder = 0
		for i := 0; i < len(source); i++ {
			accumulator := int(source[i]) + remainder*sourceBase
			digit := accumulator / targetBase
			remainder = accumulator % targetBase
			if len(quotient) > 0 || digit > 0 {
				quotient = append(quotient, digit)
			}
		}

		//ret = append(ret, byte(remainder))
		ret = append([]byte{byte(remainder)}, ret...)
		source = quotient
	}
	return ret
}

func translate(indices []byte, dictionary []byte) (ret []byte) {
	ret = make([]byte, len(indices))
	for k, v := range indices {
		ret[k] = dictionary[v]
	}
	return
}
