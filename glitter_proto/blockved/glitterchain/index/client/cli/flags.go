package cli

import (
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

func FlagSetdescription() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagDescription, "", "The description required on the dataset")
	return fs
}

func FlagSetWorkStatus() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Int32(FlagWorkStatus, 0, "The work status required on the dataset")
	return fs
}

func FlagSetHosts() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagHosts, "", "The work status required on the dataset")
	return fs
}

func FlagSetManageAddresses() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagManageAddresses, "", "The manage addresses required on the dataset")
	return fs
}

func FlagSetMeta() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.String(FlagMeta, "", "The meta required on the dataset")
	return fs
}

func FlagSetDuration() *flag.FlagSet {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Int64(FlagDuration, 1, "duration(year) required on the dataset")
	return fs
}
