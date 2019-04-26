package main

import (
	"fmt"
	"github.com/lattecake/kpaas/cmd/server/cmd"
	"github.com/spf13/cobra"
	"os"
)

var (
	httpAddr   string
	dateFormat = "2006/01/02 15:04:05"

	rootCmd = &cobra.Command{
		Use:               "kpaas-server",
		Short:             "kpaas-server服务端",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long: `kpaas平台服务端
[init,run]
`,
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "输出当前版本",
		Example: `# 输出当前版本
kpaas-server version
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(`Version: v0.1
GitRevision: 24b60e6dc7f966f8671f76050e85ef47854ecc7c
User: solacowa@gmail.com
Hub: http://github.com/lattecake/paas
GolangVersion: go1.12.4
BuildStatus: Clean`)
			return nil
		},
	}

	runCmd = &cobra.Command{
		Use:   "run",
		Short: "启动服务",
		Example: `# 启动服务
kpaas-server run -p :8080
`,
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
)

func init() {
	runCmd.PersistentFlags().StringVarP(&httpAddr, "http.addr", "p", ":8080", "启动端口, 默认: 8080")
	runCmd.PersistentFlags().StringVarP(&dateFormat, "time.format", "t", dateFormat, "日志时间格式, 默认: 2006/01/02 15:04:05")

	cmd.AddFlags(rootCmd)
	rootCmd.AddCommand(versionCmd, runCmd)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
