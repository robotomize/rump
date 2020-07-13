package codec

type Codec interface {
	Encode(v interface{}) ([]byte, error)
	Decode(bytes []byte, v interface{}) error
}
