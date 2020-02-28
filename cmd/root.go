package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"sync"

	"github.com/jonyhy96/fuck-crud/pkg/generator"
	"github.com/jonyhy96/fuck-crud/pkg/transform"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.Flags().StringP("output", "o", "", "the out put path of files")
	rootCmd.Flags().StringP("config", "c", "", "config name of config file")
	err := rootCmd.MarkFlagRequired("output")
	if err != nil {
		panic(err)
	}

	viper.SetEnvPrefix("crud")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	viper.SetConfigName("config")     // name of config file (without extension)
	viper.AddConfigPath("/etc/crud/") // path to look for the config file in
	viper.AddConfigPath("./configs")
	viper.AddConfigPath(".")
}

var (
	wg      = sync.WaitGroup{}
	rootCmd = &cobra.Command{
		Use:   "crud",
		Short: "crud generate go crud files from given config",
		Run: func(cmd *cobra.Command, args []string) {
			configName, err := cmd.Flags().GetString("config")
			if err != nil {
				panic(err)
			}
			if configName != "" {
				viper.SetConfigName(configName)
			}
			err = viper.ReadInConfig() // Find and read the config file
			if err != nil {            // Handle errors reading the config file
				panic(fmt.Errorf("Fatal error config file: %w", err))
			}
			output, _ := cmd.Flags().GetString("output")
			if _, err := os.Stat(output); os.IsNotExist(err) {
				os.MkdirAll(output, os.ModePerm)
			}
			generate(output)
		},
	}
)

func generate(outputPath string) {
	var (
		config transform.Config
		wg     = sync.WaitGroup{}
	)

	err := viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	var templates = map[string]string{ // TODO change this into read from template files by given arguments.
		generator.HandlerTpl:    "handler.go",
		generator.ServiceTpl:    "service.go",
		generator.RepositoryTpl: "repository.go",
		generator.InternalTpl:   config.Name + ".go",
	}

	wg.Add(len(templates))
	for tpl, fileName := range templates {
		go func(template string, name string) {
			defer wg.Done()
			f, err := os.Create(path.Join(outputPath, name))
			if err != nil {
				panic(err)
			}
			err = generator.Generate(config, template, f)
			if err != nil {
				panic(err)
			}
		}(tpl, fileName)
	}
	wg.Wait()
}

// Execute command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
