package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	tmjson "github.com/number571/tendermint/libs/json"
	"github.com/number571/tendermint/privval"
	"github.com/number571/tendermint/types"

	"github.com/number571/tendermint/crypto/gost256"
	"github.com/number571/tendermint/crypto/gost512"
)

// GenValidatorCmd allows the generation of a keypair for a
// validator.
var GenValidatorCmd = &cobra.Command{
	Use:     "gen-validator",
	Aliases: []string{"gen_validator"},
	Short:   "Generate new validator keypair",
	RunE:    genValidator,
	PreRun:  deprecateSnakeCase,
}

func init() {
	GenValidatorCmd.Flags().StringVar(&keyType, "key", types.ABCIPubKeyTypeGost512,
		"Key type to generate privval file with. Options: '"+gost512.KeyType+"', '"+gost256.KeyType+"'")
}

func genValidator(cmd *cobra.Command, args []string) error {
	pv, err := privval.GenFilePV("", "", keyType)
	if err != nil {
		return err
	}

	jsbz, err := tmjson.Marshal(pv)
	if err != nil {
		return fmt.Errorf("validator -> json: %w", err)
	}

	fmt.Printf(`%v
`, string(jsbz))

	return nil
}
