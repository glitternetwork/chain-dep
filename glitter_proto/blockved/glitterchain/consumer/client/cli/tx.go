package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/glitternetwork/chain-dep/glitter_proto/blockved/glitterchain/consumer/types"
	"github.com/spf13/cobra"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(GetCmdTxPledge())
	cmd.AddCommand(GetCmdTxReleasePledge())
	return cmd
}

func GetCmdTxPledge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pledge <dataset_name> <amount> ",
		Short: "pledge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var datasetName = args[0]
			amount, _ := sdk.NewIntFromString(args[1])
			msg := types.NewPledgeRequest(fromAddress, datasetName, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdTxReleasePledge() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "release_pledge <dataset_name> <amount> ",
		Short: "release_pledge",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var datasetName = args[0]
			amount, _ := sdk.NewIntFromString(args[1])
			msg := types.NewReleasePledgeRequest(fromAddress, datasetName, amount)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
