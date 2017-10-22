package encrypted

import (
  "encoding/base64"
)

type Secret struct {
  Data string
  Nonce string
  Salt string
}

func Create(data []byte, nonce [24]byte, salt []byte) Secret {
  return Secret{Data: base64.StdEncoding.EncodeToString(data), Nonce: base64.StdEncoding.EncodeToString(nonce[:]), Salt: base64.StdEncoding.EncodeToString(salt)}
}
