package encoding

type (
	Encoder func(v any) ([]byte, error)
	Decoder func(data []byte, v any) error
)

type Codec interface {
	Encode(v any) ([]byte, error)
	Decode(data []byte, v any) error
}

type codec struct {
	encoder Encoder
	decoder Decoder
}

func (c *codec) Encode(v any) ([]byte, error) {
	return c.encoder(v)
}

func (c *codec) Decode(data []byte, v any) error {
	return c.decoder(data, v)
}

func NewCodec(enc Encoder, dec Decoder) Codec {
	return &codec{encoder: enc, decoder: dec}
}
