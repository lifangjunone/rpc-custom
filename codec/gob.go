package codec

import (
	"bytes"
	"encoding/gob"
)

// GobEncode object -> gob -> []byte
func GobEncode(obj interface{}) ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(obj); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// GobDecode []byte -> gob -> object
func GobDecode(data []byte, obj interface{}) error {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
