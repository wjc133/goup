package commands

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	homedir string
	logger  *logrus.Logger

	ProfileFiles []string

	rootCmdVerboseFlag bool
)

func init() {
	logger = logrus.New()

	var err error
	homedir, err = os.UserHomeDir()
	if err != nil {
		logger.Fatal(err)
	}

	ProfileFiles = []string{
		filepath.Join(homedir, ".profile"),
		filepath.Join(homedir, ".zprofile"),
		filepath.Join(homedir, ".bash_profile"),
	}
}

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "goup",
		Short:             "The Go installer",
		PersistentPreRunE: preRunRoot,
	}

	rootCmd.PersistentFlags().BoolVarP(&rootCmdVerboseFlag, "verbose", "v", false, "Verbose")

	rootCmd.AddCommand(installCmd())
	rootCmd.AddCommand(setCmd())
	rootCmd.AddCommand(removeCmd())
	rootCmd.AddCommand(initCmd())
	rootCmd.AddCommand(listCmd())
	rootCmd.AddCommand(searchCmd())
	rootCmd.AddCommand(versionCmd())
	rootCmd.AddCommand(upgradeCmd())

	return rootCmd
}

func GoupBinDir() string {
	return GoBaseDir("bin")
}

func GoupCurrentDir() string {
	return GoBaseDir("current")
}

func GoupEnvFile() string {
	return GoBaseDir("env")
}

func GoupCurrentBinDir() string {
	return GoBaseDir("current", "bin")
}

func goupVersionDir(ver string) string {
	return GoBaseDir(ver)
}

func GoBaseDir(paths ...string) string {
	elem := []string{homedir, "go"}
	elem = append(elem, paths...)

	return filepath.Join(elem...)
}

func HomebrewGoDir() string {
	return "/opt/homebrew/Cellar/go"
}

func preRunRoot(cmd *cobra.Command, args []string) error {
	if rootCmdVerboseFlag {
		logger.SetLevel(logrus.DebugLevel)
	}

	return nil
}
