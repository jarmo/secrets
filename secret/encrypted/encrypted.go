package encrypted

import (
  "encoding/hex"
)

type Secret struct {
  Data string
  Nonce string
  Salt string
}

func Create(data []byte, nonce [24]byte, salt []byte) Secret {
  return Secret{Data: hex.EncodeToString(data), Nonce: hex.EncodeToString(nonce[:]), Salt: hex.EncodeToString(salt)}
}
