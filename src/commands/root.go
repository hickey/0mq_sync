package commands

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var (
  config string
  daemon string
  version bool

  OctopusCmd = &cobra.Command {
    Use: "",
    Long: "",
    Short: "",

    PersistentPreRun: func(cmd *cobra.Command, args []string) {
      // if --config is passed, attempt to parse the config file
      if config != "" {

        // get the filepath
        abs, err := filepath.Abs(config)
        if err != nil {
          lumber.Error("Error reading filepath: ", err.Error())
        }

        // get the config name
        base := filepath.Base(abs)

        // get the path
        path := filepath.Dir(abs)

        //
        viper.SetConfigName(strings.Split(base, ".")[0])
        viper.AddConfigPath(path)

        // Find and read the config file; Handle errors reading the config file
        if err := viper.ReadInConfig(); err != nil {
          lumber.Fatal("Failed to read config file: ", err.Error())
          os.Exit(1)
        }
      }
    },

    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("You said hello")
    },
  }
)


func init() {

  // set config defaults
  viper.SetDefault("garbage-collect", false)

  // persistent flags
  RootCmd.PersistentFlags().StringP("backend", "b", "file://", "Hoarder backend driver")
  RootCmd.PersistentFlags().Uint64P("clean-after", "g", uint64(time.Now().Unix()), "Age, in seconds, after which data is deemed garbage")
  RootCmd.PersistentFlags().StringP("host", "H", "127.0.0.1", "Hoarder hostname/IP")
  RootCmd.PersistentFlags().BoolP("insecure", "i", true, "Whether or not to start the Hoarder server with TLS")
  RootCmd.PersistentFlags().String("log-type", "stdout", "The type of logging (stdout, file)")
  RootCmd.PersistentFlags().String("log-file", "/var/log/hoarder.log", "If log-type=file, the /path/to/logfile; ignored otherwise")
  RootCmd.PersistentFlags().String("log-level", "INFO", "Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL)")
  RootCmd.PersistentFlags().StringP("port", "p", "7410", "Hoarder port")
  RootCmd.PersistentFlags().StringP("token", "t", "", "Auth token used when connecting to a secure Hoarder")

  //
  viper.BindPFlag("backend", RootCmd.PersistentFlags().Lookup("backend"))
  viper.BindPFlag("clean-after", RootCmd.PersistentFlags().Lookup("clean-after"))
  viper.BindPFlag("host", RootCmd.PersistentFlags().Lookup("host"))
  viper.BindPFlag("insecure", RootCmd.PersistentFlags().Lookup("insecure"))
  viper.BindPFlag("log-type", RootCmd.PersistentFlags().Lookup("log-type"))
  viper.BindPFlag("log-file", RootCmd.PersistentFlags().Lookup("log-file"))
  viper.BindPFlag("log-level", RootCmd.PersistentFlags().Lookup("log-level"))
  viper.BindPFlag("port", RootCmd.PersistentFlags().Lookup("port"))
  viper.BindPFlag("token", RootCmd.PersistentFlags().Lookup("token"))

  // local flags;
  RootCmd.Flags().StringVar(&config, "config", "", "/path/to/config.yml")
  RootCmd.Flags().BoolVar(&daemon, "server", false, "Run hoarder as a server")
  RootCmd.Flags().BoolVarP(&version, "version", "v", false, "Display the current version of this CLI")

  // commands
  RootCmd.AddCommand(addCmd)
  RootCmd.AddCommand(listCmd)
  RootCmd.AddCommand(removeCmd)
  RootCmd.AddCommand(showCmd)
  RootCmd.AddCommand(updateCmd)

  // hidden/aliased commands
  RootCmd.AddCommand(createCmd)
  RootCmd.AddCommand(deleteCmd)
  RootCmd.AddCommand(destroyCmd)
  RootCmd.AddCommand(fetchCmd)
  RootCmd.AddCommand(getCmd)
}