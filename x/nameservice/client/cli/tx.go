package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/manyaagarwal/nameservice/x/nameservice/types"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func GetTxCmd (storeKey string, cdc *codec.Codec) *cobra.Command {
	nameserviceTxCmd := &cobra.Command {
		Use: types.ModuleName,
		Short: "Nameservice transaction subcommands",
		DisableFlagParsing: true,
		SuggestionsMinimumDistance: 2,
		RunE: client.ValidateCmd,
	}
	nameserviceTxCmd.AddCommand(client.PostCommands(
		GetCmdBuyName(cdc),
		GetCmdSetName(cdc),
	)...)
	return nameserviceTxCmd
}

func GetCmdBuyName(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command {
		USe: "buy-name [name] [account]",
		Short: "bid for existing name or claim new name",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			msg := types.NewMsgBuyName(args[0], coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}
			cliCtx.PrintResponse = true
			return utils.GenerateOrBoradcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdSetName (cdc *codec.Codec) *cobra.Command {
	return &cobra.command {
		Use: "set-name [name] [value]",
		Short: "set the value associated with a name that you own",
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			msg := types.NewMsgSetName(args[0], args[1], cliCtx.GetFromAccAddress())

			err := msg.ValidateBasic()

			if err != nil { return err }

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsg(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}