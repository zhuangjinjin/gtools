package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/zhuangjinjin/gtools/pkg/version"
	"github.com/zhuangjinjin/gtools/util"
)

var idCardStr string

var rootCmd = &cobra.Command{
	Use: "gtools",
	Short: "golang common biz tools",
	Long: "a golang common biz tools, includes idcard verify ...",
	Version: version.Get().ToString(),
	Run: func(cmd *cobra.Command, args []string) {
		if idCardStr != "" {
			verifyIdcard(idCardStr)
		}
	}, 
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.Flags().StringVar(&idCardStr, "verify-idcard", "", "idcard, like 350521XXXXXXXXXXX8.")
}

func verifyIdcard(idCardStr string) {
	result := util.VerifyIdCard(idCardStr)
	fmt.Printf("idcard: %s, verify result: %t\n", idCardStr, result)
}