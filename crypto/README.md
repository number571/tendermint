# crypto

crypto is the cryptographic package adapted for Tendermint's uses

## Importing it

To get the interfaces,
`import "github.com/number571/tendermint/crypto"`

For any specific algorithm, use its specific module e.g.
`import "github.com/number571/tendermint/crypto/gost512"`

## Binary encoding

For Binary encoding, please refer to the [Tendermint encoding specification](https://docs.tendermint.com/master/spec/blockchain/encoding.html).

## JSON Encoding

JSON encoding is done using tendermint's internal json encoder. For more information on JSON encoding, please refer to [Tendermint JSON encoding](https://github.com/number571/tendermint/blob/ccc990498df70f5a3df06d22476c9bb83812cbe3/libs/json/doc.go)

```go
Example JSON encodings:

gost512.PrivKey     - {"type":"tendermint/PrivKey512","value":"UTEwYjdhYTBkNTYwOTM3NDI5OWQwYjc0YWQyN2ZlODEyMDVkNDRmNWJlYTY3ZTUwNjY4ZTZlMmQzZTBmZTMyMGM3Yjc5N2I2NDRjZDA4Y2FjNTliNDk5NDA3YTMxYzIxODM4NjRjYmMwNTkzMmJkZDA0MjNiMzc5ZWRjOTg0OGNl"}
gost512.PubKey      - {"type":"tendermint/PubKey512","value":"UQYgAAA9LgAATUFHMQAEAAAwFQYJKoUDBwECAQIBBggqhQMHAQECA8OWSvhanUIAIQJNpbhLA3qbO7q8wu5Bcvu9f/4Grxee3yqSaPSA9sMne6auI2P9bLYagEdJDMZyTYIWGj/RUDhOyWIdymUWDwA6EiiGFqqrBK5uoq+X0xxmrT5ikC1VbGueC0EaZLYqqQpaG37Fga1vFYPPz4qIYLi3in6o5CYb"}
gost256.PubKey   - {"type":"tendermint/PrivKey256","value":"UDMxNzU0Mjk3Y2FjYTdiMmJjYzkxYWM4NTFjNWUzOTJjZWU0MDFhN2ZkNjQzMDFhODI2OGUxM2E0NjBlNmE2MDM3Yjc5N2I2NDRjZDA4Y2FjNTliNDk5NDA3YTMxYzIxODM4NjRjYmMwNTkzMmJkZDA0MjNiMzc5ZWRjOTg0OGNl"}
gost256.PubKey    - {"type":"tendermint/PubKey256","value":"UAYgAABJLgAATUFHMQACAAAwEwYHKoUDAgIjAQYIKoUDBwEBAgJaWews1GShsVvOwilvu3LrO6rBhdyVuBvZjUmaY+oOS9+xTff8DtlpgaukttrSS+q1L0IGC5iXP0K5q1vwTs9m"}
```
