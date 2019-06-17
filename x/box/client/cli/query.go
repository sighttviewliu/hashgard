package cli

import (
	"strings"

	"github.com/hashgard/hashgard/x/box/config"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	boxqueriers "github.com/hashgard/hashgard/x/box/client/queriers"
	"github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Box query
// QueryCmd implements the query box command.
func QueryCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "box [box-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query the details of the account box",
		Long:    "Query the details of the account box",
		Example: "$ hashgardcli bank box boxab3jlxpt2ps",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processQueryBoxCmd(cdc, args[0])
		},
	}
}

// ProcessQueryBoxParamsCmd implements the query box params command.
func ProcessQueryBoxParamsCmd(cdc *codec.Codec, boxType string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	res, err := boxqueriers.QueryBoxParams(cliCtx)
	if err != nil {
		return err
	}
	var params config.Params
	cdc.MustUnmarshalJSON(res, &params)
	return cliCtx.PrintOutput(utils.GetBoxParams(params, boxType))
}

// ProcessQueryBoxCmd implements the query box command.
func ProcessQueryBoxCmd(cdc *codec.Codec, boxType string, id string) error {
	if boxutils.GetBoxTypeByValue(id) != boxType {
		return errors.Errorf(errors.ErrUnknownBox(id))
	}
	return processQueryBoxCmd(cdc, id)
}

// ProcessQueryBoxCmd implements the query box command.
func processQueryBoxCmd(cdc *codec.Codec, id string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	if err := boxutils.CheckId(id); err != nil {
		return errors.Errorf(err)
	}
	// Query the box
	res, err := boxqueriers.QueryBoxByID(id, cliCtx)
	if err != nil {
		return err
	}
	var box types.BoxInfo
	cdc.MustUnmarshalJSON(res, &box)
	return cliCtx.PrintOutput(utils.GetBoxInfo(box))
}

// ProcessListBoxCmd implements the query box command.
func ProcessListBoxCmd(cdc *codec.Codec, boxType string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	_, ok := types.BoxType[boxType]
	if !ok {
		return errors.Errorf(errors.ErrUnknownBoxType())
	}
	address, err := sdk.AccAddressFromBech32(viper.GetString(FlagAddress))
	if err != nil {
		return err
	}
	boxQueryParams := params.BoxQueryParams{
		StartId: viper.GetString(FlagStartId),
		BoxType: boxType,
		Owner:   address,
		Limit:   viper.GetInt(FlagLimit),
	}
	// Query the box
	res, err := boxqueriers.QueryBoxsList(boxQueryParams, cdc, cliCtx)
	if err != nil {
		return err
	}
	var boxs types.BoxInfos
	cdc.MustUnmarshalJSON(res, &boxs)
	return cliCtx.PrintOutput(utils.GetBoxList(boxs, boxQueryParams.BoxType))
}

// ProcessSearchBoxsCmd implements the query box command.
func ProcessSearchBoxsCmd(cdc *codec.Codec, boxType string, name string) error {
	_, ok := types.BoxType[boxType]

	if !ok {
		return errors.Errorf(errors.ErrUnknownBoxType())
	}
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	// Query the box
	res, err := boxqueriers.QueryBoxByName(boxType, strings.ToLower(name), cliCtx)
	if err != nil {
		return err
	}
	var boxs types.BoxInfos
	cdc.MustUnmarshalJSON(res, &boxs)
	return cliCtx.PrintOutput(utils.GetBoxList(boxs, boxType))
}
