package hystrix

import "github.com/spf13/cobra"

var cfg string

func NewHystrixCommand() *cobra.Command {
	cmd := &cobra.Command{}

	return cmd
}
