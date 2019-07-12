package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Verbose flag
var (
	Verbose        bool
	cfgFile        string
	namespace      string
	basepath       string
	connectTimeout int
	readTimeout    int
	writeTimeout   int
	retries        int
	protocols      []string
	methods        []string
	host           string
	upstreamHost   string
	upstreamPort   int
)

var rootCmd = &cobra.Command{
	Use:   "api-gateway",
	Short: "api-gateway is a generic vendor agnostic tool that manage and provision API Gateways",
	Long: `An agnostic cli tool that manage any API Gateway.
			Built with love and available at https://github.com/Amzani/api-gateway-cli
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

//Execute function
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.EnableCommandSorting = false

	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.api-blueprint.yaml)")
	rootCmd.PersistentFlags().StringVar(&namespace, "namespace", "", "Your k8s namespace")
	viper.BindPFlag("namespace", rootCmd.PersistentFlags().Lookup("namespace"))

	rootCmd.PersistentFlags().StringVar(&basepath, "basepath", "", "Your API Base path")
	viper.BindPFlag("basepath", rootCmd.PersistentFlags().Lookup("basepath"))

	rootCmd.PersistentFlags().IntVar(&connectTimeout, "connect_timeout", 5000, "upstream connect timeout")
	viper.BindPFlag("connect_timeout", rootCmd.PersistentFlags().Lookup("connect_timeout"))

	rootCmd.PersistentFlags().IntVar(&readTimeout, "read_timeout", 5000, "upstream read timeout")
	viper.BindPFlag("read_timeout", rootCmd.PersistentFlags().Lookup("read_timeout"))

	rootCmd.PersistentFlags().IntVar(&writeTimeout, "write_timeout", 5000, "upstream write timeout")
	viper.BindPFlag("write_timeout", rootCmd.PersistentFlags().Lookup("write_timeout"))

	rootCmd.PersistentFlags().IntVar(&retries, "retries", 5, "number of retries")
	viper.BindPFlag("retries", rootCmd.PersistentFlags().Lookup("retries"))

	rootCmd.PersistentFlags().StringArrayVarP(&protocols, "protocols", "p", []string{}, "supported api protocols")
	viper.BindPFlag("protocols", rootCmd.PersistentFlags().Lookup("protocols"))

	rootCmd.PersistentFlags().StringArrayVarP(&methods, "methods", "m", []string{}, "supported api methods")
	viper.BindPFlag("methods", rootCmd.PersistentFlags().Lookup("methods"))

	rootCmd.PersistentFlags().StringVar(&host, "host", "", "api host")
	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))

	rootCmd.PersistentFlags().StringVar(&upstreamHost, "upstream_host", "", "upstream")
	viper.BindPFlag("upstream_host", rootCmd.PersistentFlags().Lookup("upstream_host"))

	rootCmd.PersistentFlags().IntVar(&upstreamPort, "upstream_port", 80, "upstream port")
	viper.BindPFlag("upstream_port", rootCmd.PersistentFlags().Lookup("upstream_port"))

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.MarkFlagRequired("namespace")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".api-gateway" (without extension).
		viper.AddConfigPath(dir)
		viper.SetConfigName("api-blueprint")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
