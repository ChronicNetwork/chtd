package main

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	chtCli "github.com/ChronicNetwork/cht/x/cht/client/cli"
)

func AddGenesisChtMsgCmd(defaultNodeHome string) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "add-cht-genesis-message",
		Short:                      "CHT genesis subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	genesisIO := chtCli.NewDefaultGenesisIO()
	txCmd.AddCommand(
		chtCli.GenesisStoreCodeCmd(defaultNodeHome, genesisIO),
		chtCli.GenesisInstantiateContractCmd(defaultNodeHome, genesisIO),
		chtCli.GenesisExecuteContractCmd(defaultNodeHome, genesisIO),
		chtCli.GenesisListContractsCmd(defaultNodeHome, genesisIO),
		chtCli.GenesisListCodesCmd(defaultNodeHome, genesisIO),
	)
	return txCmd
}
