package cmd

import (
	"n3dr/cli"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	names  bool
	count  bool
	backup bool
)

// repositoriesCmd represents the repositories command
var repositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Count the number of repositories or return their names",
	Long: `Count the number of repositories, count the total or
download artifacts from all repositories`,
	Run: func(cmd *cobra.Command, args []string) {
		if !(names || count || backup) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
			log.Fatal("One of the required flags \"names\", \"count\" or \"backup\" not set")
		}
		n := cli.Nexus3{URL: n3drURL, User: n3drUser, Pass: n3drPass, APIVersion: apiVersion, ZIP: zip, ZipName: zipName, DownloadDirName: downloadDirName}
		if err := n.ValidateNexusURL(); err != nil {
			log.Fatal(err)
		}
		if names {
			if err := n.RepositoryNames(); err != nil {
				log.Fatal(err)
			}
		}
		if count {
			if err := n.CountRepositories(); err != nil {
				log.Fatal(err)
			}
		}
		if backup {
			err := n.Downloads(regex)
			if err != nil {
				log.Fatal(err)
			}
		}
	},
	Version: rootCmd.Version,
}

func init() {
	repositoriesCmd.Flags().BoolVarP(&names, "names", "a", false, "print all repository names")
	repositoriesCmd.Flags().BoolVarP(&count, "count", "c", false, "count the number of repositories")
	repositoriesCmd.Flags().BoolVarP(&backup, "backup", "b", false, "backup artifacts from all repositories")
	repositoriesCmd.Flags().StringVarP(&regex, "regex", "x", ".*", "only download artifacts that match a regular expression, e.g. 'some/group42'")

	rootCmd.AddCommand(repositoriesCmd)
}
