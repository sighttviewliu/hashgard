package cli

import (
	"github.com/cosmos/cosmos-sdk/codec"
	boxcli "github.com/hashgard/hashgard/x/box/client/cli"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/spf13/cobra"
)

func GetDescriptionCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "describe [id] [description-file]",
		Args:    cobra.ExactArgs(2),
		Short:   "Describe the lock",
		Long:    "Owner can set description of the lock, and the description need to be in json format. You can customize preferences or use recommended templates.",
		Example: "$ hashgardcli lock describe boxab3jlxpt2ps path/description.json --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return boxcli.ProcessBoxDescriptionCmd(cdc, types.Lock, args[0], args[1])
		},
	}
	return cmd
}
