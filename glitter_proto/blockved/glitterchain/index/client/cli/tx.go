package cli

import (
	"errors"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/glitternetwork/chain-dep/glitter_proto/blockved/glitterchain/index/types"
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

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(GetCmdTxSQLExec())
	cmd.AddCommand(GetCmdTxSQLGrant())
	cmd.AddCommand(GetCmdTxBindHost())
	cmd.AddCommand(GetCmdTxCreateDataset())
	cmd.AddCommand(GetCmdTxEditDataset())
	cmd.AddCommand(GetCmdTxEditTable())
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

// GetCmdTxSQLExec returns the command that broadcasts to sql-grant
func GetCmdTxBindHost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bind-host <database> <url>",
		Short: "grant permission",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			ownerAddress := clientCtx.GetFromAddress()
			var (
				database, url string
			)
			database, url = args[0], args[1]
			if len(url) > 1000 {
				return errors.New("url is too long")
			}
			msg := types.NewBindHostRequest(ownerAddress, database, url)
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
func GetCmdTxCreateDataset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-dataset <dataset_name>",
		Short: "create dataset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var (
				datasetName string
			)
			datasetName = args[0]
			if err != nil {
				return err
			}
			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			hosts, err := cmd.Flags().GetString(FlagHosts)
			if err != nil {
				return err
			}

			manageAddresses, err := cmd.Flags().GetString(FlagManageAddresses)
			if err != nil {
				return err
			}

			workstatus, err := cmd.Flags().GetInt32(FlagWorkStatus)
			if err != nil {
				return err
			}

			msg := types.NewCreateDatasetRequest(
				fromAddress,
				datasetName,
				types.ServiceStatus(workstatus),
				hosts,
				manageAddresses,
				description,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetdescription())
	cmd.Flags().AddFlagSet(FlagSetWorkStatus())
	cmd.Flags().AddFlagSet(FlagSetHosts())
	cmd.Flags().AddFlagSet(FlagSetManageAddresses())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdTxSQLExec returns the command that broadcasts to sql-grant
func GetCmdTxEditDataset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-dataset <dataset_name>",
		Short: "edit dataset",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var (
				datasetName string
			)
			datasetName = args[0]
			if err != nil {
				return err
			}
			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			hosts, err := cmd.Flags().GetString(FlagHosts)
			if err != nil {
				return err
			}

			manageAddresses, err := cmd.Flags().GetString(FlagManageAddresses)
			if err != nil {
				return err
			}

			workstatus, err := cmd.Flags().GetInt32(FlagWorkStatus)
			if err != nil {
				return err
			}

			msg := types.NewEditDatasetRequest(
				fromAddress,
				datasetName,
				types.ServiceStatus(workstatus),
				hosts,
				manageAddresses,
				description,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().AddFlagSet(FlagSetdescription())
	cmd.Flags().AddFlagSet(FlagSetWorkStatus())
	cmd.Flags().AddFlagSet(FlagSetHosts())
	cmd.Flags().AddFlagSet(FlagSetManageAddresses())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdTxSQLExec returns the command that broadcasts to sql-grant
func GetCmdTxEditTable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-table <dataset_name> <table_name> <meta>",
		Short: "edit table",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var (
				datasetName, tableName, meta string
			)
			datasetName, tableName, meta = args[0], args[1], args[2]

			msg := types.NewEditTableRequest(fromAddress, datasetName, tableName, meta)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
