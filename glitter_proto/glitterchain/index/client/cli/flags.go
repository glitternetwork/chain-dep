package cli

import (
	"github.com/glitternetwork/chain-dep/utils"
	flag "github.com/spf13/pflag"
)

const (
	FlagDescription     = "description"
	FlagWorkStatus      = "work-status"
	FlagHosts           = "hosts"
	FlagManageAddresses = "manage-addresses"
	FlagMeta            = "meta"
	FlagDuration        = "duration"
)

func FlagSetDescription() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagDescription, utils.DefaultStringField, "The description required on the dataset")
	return fs
}

func FlagSetWorkStatus() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Int64(FlagWorkStatus, 0, "The work status required on the dataset")
	return fs
}

func FlagSetHosts() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagHosts, utils.DefaultStringField, "The work status required on the dataset")
	return fs
}

func FlagSetManageAddresses() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagManageAddresses, utils.DefaultStringField, "The manage addresses required on the dataset")
	return fs
}

func FlagSetMeta() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagMeta, utils.DefaultStringField, "The meta required on the dataset")
	return fs
}

func FlagSetDuration() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Int64(FlagDuration, 0, "duration(second) required on the dataset")
	return fs
}
