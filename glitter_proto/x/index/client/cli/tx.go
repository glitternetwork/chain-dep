package cli

import (
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/glitternetwork/chain-dep/glitter_proto/x/index/types"
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

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(GetCmdTxSQLExec())
	cmd.AddCommand(GetCmdTxSQLGrant())
	return cmd
}

// GetCmdTxSQLExec returns the command that broadcasts to sql-exec
func GetCmdTxSQLExec() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sql-exec <sql_statement>",
		Short: "execute sql",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			ownerAddress := clientCtx.GetFromAddress()
			if len(args[0]) == 0 {
				return errors.New("args error: empty sql")
			}
			sql := args[0]
			msg := types.NewSQLExecRequest(ownerAddress, sql, nil)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdTxSQLExec returns the command that broadcasts to sql-grant
func GetCmdTxSQLGrant() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sql-grant <on_database> <on_table> <to_uid> <role>",
		Short: "grant permission",
		Args:  cobra.RangeArgs(3, 4),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			ownerAddress := clientCtx.GetFromAddress()
			var (
				onDatabase, onTable, toUID, role string
			)
			grantDB := true
			if len(args) == 3 {
				onDatabase, toUID, role = args[0], args[1], args[2]
				grantDB = true
			} else {
				onDatabase, onTable, toUID, role = args[0], args[1], args[2], args[3]
				grantDB = false
			}

			if len(onDatabase) == 0 {
				return errors.New("args error: empty database")
			}
			if !grantDB && len(onTable) == 0 {
				return errors.New("args error: empty table")
			}
			if len(toUID) == 0 {
				return errors.New("args error: empty to_uid")
			}
			if len(role) == 0 {
				return errors.New("args error: empty role")
			}
			msg := types.NewSQLGrantRequest(ownerAddress, onDatabase, onTable, toUID, role)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
