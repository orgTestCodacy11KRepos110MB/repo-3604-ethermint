package client

import (
	"github.com/evmos/ethermint/client/eip712"
	"github.com/spf13/cobra"
)

func EIP712Commands() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eip-712",
		Short: "EIP-712 metatransactions utility commands",
	}
	cmd.AddCommand(eip712.DataTypeCommand())
	return cmd
}