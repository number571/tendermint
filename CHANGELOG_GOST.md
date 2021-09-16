
## ЗАМЕНА КРИПТОГРАФИЧЕСКИХ ФУНКЦИЙ TENDERMINT НА ГОСТ ФУНКЦИИ

1.  Изменён файл "crypto/batch/batch.go";
2.  Изменён файл "crypto/encoding/codec.go";
3.  Изменены файлы в директории "crypto/ed25519";
4.  Переименован пакет "crypto/ed25519" на "crypto/gost512";
5.  Изменены файлы в директории "crypto/secp256k1";
6.  Переименован пакет "crypto/secp256k1" на "crypto/gost256";
7.  Удалена директория "crypto/sr25519";
8.  Удалена директория "crypto/xchacha20poly1305";
9.  Удалена директория "crypto/xsalsa20symmetric";
10. Изменены файлы в директории "crypto/merkle";
11. Изменены файлы в директории "crypto/tmhash";
12. Измененён файл "crypto/hash.go";
13. Измененён файл "crypto/random.go";
14. Измененён файл "internal/p2p/conn/secret_connection.go";
15. Измененён файл "privval/secret_connection.go";

## ЗАВИСИМОСТИ

1.  Измененён файл "abci/types/pubkey.go";
2.  Измененён файл "cmd/tendermint/commands/gen_node_key.go";
3.  Измененён файл "cmd/tendermint/commands/gen_validator.go";
4.  Измененён файл "cmd/tendermint/commands/init.go";
5.  Измененён файл "cmd/tendermint/commands/reset_priv_validator.go";
6.  Измененён файл "cmd/tendermint/commands/testnet.go";
7.  Измененён файл "privval/file.go";
8.  Измененён файл "privval/socket_listeners.go";
9.  Измененён файл "privval/utils.go";
10. Измененён файл "proto/tendermint/crypto/keys.pb.go" (НЕОБХОДИМО ПЕРЕГЕНЕРИРОВАТЬ!)
11. Измененён файл "proto/tendermint/crypto/keys.proto" (НЕОБХОДИМО ПЕРЕГЕНЕРИРОВАТЬ!)
12. Измененён файл "types/node_key.go" 
13. Измененён файл "types/params.go"
14. Измененён файл "types/priv_validator.go"
15. Измененён файл "types/signable.go"
16. Измененён файл "internal/statesync/snapshots.go";

## ПРОМЕЖУТОЧНЫЕ, ДОПОЛНИТЕЛЬНЫЕ ИЗМЕНЕНИЯ

1. Изменены ссылки "github.com/tendermint/tendermint" на "github.com/number571/tendermint";
2. Изменён файл "config/config.go": добавлен параметр "DisableLegacy: true" в функцию DefaultP2PConfig;
