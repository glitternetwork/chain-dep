package cli

import (
	"fmt"
	"github.com/glitternetwork/chain-dep/utils"
	"github.com/pkg/errors"
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

	cmd.AddCommand(GetCmdTxCreateDataset())
	cmd.AddCommand(GetCmdTxEditDataset())
	cmd.AddCommand(GetCmdTxEditTable())
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
			if datasetName == "" {
				return errors.New("param dataset name error")
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

			workstatus, err := cmd.Flags().GetInt64(FlagWorkStatus)
			if err != nil {
				return err
			}
			if workstatus == utils.DefaultInt64FieldUnsetdata {
				return errors.New("param work status error")
			}

			meta, err := cmd.Flags().GetString(FlagMeta)
			if err != nil {
				return err
			}
			if meta == utils.DefaultStringFieldUnsetData {
				return errors.New("param meta error")
			}

			duration, err := cmd.Flags().GetInt64(FlagDuration)
			if err != nil {
				return err
			}
			if duration == utils.DefaultInt64FieldUnsetdata {
				return errors.New("unset param duration")
			}

			if duration <= 0 {
				return errors.New("param useYear must greater than 0")
			}

			msg := types.NewCreateDatasetRequest(
				fromAddress,
				datasetName,
				types.ServiceStatus(workstatus),
				hosts,
				meta,
				manageAddresses,
				description,
				duration,
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
	cmd.Flags().AddFlagSet(FlagSetMeta())
	cmd.Flags().AddFlagSet(FlagSetDuration())
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

			workstatus, err := cmd.Flags().GetInt64(FlagWorkStatus)
			if err != nil {
				return err
			}

			meta, err := cmd.Flags().GetString(FlagMeta)
			if err != nil {
				return err
			}
			msg := types.NewEditDatasetRequest(
				fromAddress,
				datasetName,
				types.ServiceStatus(workstatus),
				hosts,
				manageAddresses,
				meta,
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
