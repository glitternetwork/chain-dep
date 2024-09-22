package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/glitternetwork/chain-dep/glitter_proto/glitterchain/index/types"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"strconv"
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

	cmd.AddCommand(GetCmdTxCreateDataset())
	cmd.AddCommand(GetCmdTxEditDataset())
	cmd.AddCommand(GetCmdTxEditTable())
	cmd.AddCommand(GetCmdTxRenewalDataset())
	return cmd
}

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
			var isFindWorkStatus bool
			for k, _ := range types.ServiceStatus_name {
				if k == int32(workstatus) {
					isFindWorkStatus = true
					break
				}
			}
			if !isFindWorkStatus {
				return errors.New("param work status error")
			}
			duration, err := cmd.Flags().GetInt64(FlagDuration)
			if err != nil {
				return err
			}
			if duration <= 0 {
				return errors.New("param duration must greater than 0")
			}
			msg := types.NewCreateDatasetRequest(
				fromAddress,
				datasetName,
				types.ServiceStatus(workstatus),
				hosts,
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

	cmd.Flags().AddFlagSet(FlagSetWorkStatus())
	cmd.Flags().AddFlagSet(FlagSetHosts())
	cmd.Flags().AddFlagSet(FlagSetManageAddresses())
	cmd.Flags().AddFlagSet(FlagSetDuration())
	cmd.Flags().AddFlagSet(FlagSetDescription())
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

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

	cmd.Flags().AddFlagSet(FlagSetDescription())
	cmd.Flags().AddFlagSet(FlagSetWorkStatus())
	cmd.Flags().AddFlagSet(FlagSetHosts())
	cmd.Flags().AddFlagSet(FlagSetManageAddresses())
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdTxEditTable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit-table <dataset_name> <table_name> <description>",
		Short: "edit table",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var (
				datasetName, tableName, description string
			)
			datasetName, tableName, description = args[0], args[1], args[3]

			msg := types.NewEditTableRequest(fromAddress, datasetName, tableName, description)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func GetCmdTxRenewalDataset() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "renewal-dataset <dataset_name>",
		Short: "renewal dataset",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			fromAddress := clientCtx.GetFromAddress()
			var datasetName = args[0]
			duration, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewRenewalDatasetRequest(fromAddress, datasetName, duration)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
