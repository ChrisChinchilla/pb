// Copyright (c) 2024 Parseable, Inc
//
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"errors"
	"fmt"
	"os"
	"sync"

	pb "pb/cmd"
	"pb/pkg/analytics"
	"pb/pkg/config"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

// populated at build time
var (
	Version string
	Commit  string
)

var (
	versionFlag      = "version"
	versionFlagShort = "v"
)

func defaultInitialProfile() config.Profile {
	return config.Profile{
		URL:      "https://demo.parseable.com",
		Username: "admin",
		Password: "admin",
	}
}

// Root command
var cli = &cobra.Command{
	Use:               "pb",
	Short:             "\nParseable command line interface",
	Long:              "\npb is the command line interface for Parseable",
	PersistentPreRunE: analytics.CheckAndCreateULID,
	RunE: func(command *cobra.Command, _ []string) error {
		if p, _ := command.Flags().GetBool(versionFlag); p {
			pb.PrintVersion(Version, Commit)
			return nil
		}
		return errors.New("no command or flag supplied")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "cli", args)
		}()
	},
}

var profile = &cobra.Command{
	Use:               "profile",
	Short:             "Manage different Parseable targets",
	Long:              "\nuse profile command to configure different Parseable instances. Each profile takes a URL and credentials.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "profile", args)
		}()
	},
}

var schema = &cobra.Command{
	Use:   "schema",
	Short: "Generate or create schemas for JSON data or Parseable streams",
	Long: `The "schema" command allows you to either:
  - Generate a schema automatically from a JSON file for analysis or integration.
  - Create a custom schema for Parseable streams (PB streams) to structure and process your data.

Examples:
  - To generate a schema from a JSON file:
      pb schema generate --file=data.json
  - To create a schema for a PB stream:
      pb schema create --stream-name=my_stream --config=data.json
`,
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "generate", args)
		}()
	},
}

var user = &cobra.Command{
	Use:               "user",
	Short:             "Manage users",
	Long:              "\nuser command is used to manage users.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "user", args)
		}()
	},
}

var role = &cobra.Command{
	Use:               "role",
	Short:             "Manage roles",
	Long:              "\nrole command is used to manage roles.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "role", args)
		}()
	},
}

var stream = &cobra.Command{
	Use:               "stream",
	Short:             "Manage streams",
	Long:              "\nstream command is used to manage streams.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "stream", args)
		}()
	},
}

var query = &cobra.Command{
	Use:               "query",
	Short:             "Run SQL query on a log stream",
	Long:              "\nRun SQL query on a log stream. Default output format is json. Use -i flag to open interactive table view.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "query", args)
		}()
	},
}

var cluster = &cobra.Command{
	Use:               "cluster",
	Short:             "Cluster operations for Parseable.",
	Long:              "\nCluster operations for Parseable cluster on Kubernetes.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "install", args)
		}()
	},
}

var list = &cobra.Command{
	Use:               "list",
	Short:             "List parseable on kubernetes cluster",
	Long:              "\nlist command is used to list Parseable oss installations.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "install", args)
		}()
	},
}

var show = &cobra.Command{
	Use:               "show",
	Short:             "Show outputs values defined when installing Parseable on kubernetes cluster",
	Long:              "\nshow command is used to get values in Parseable.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "install", args)
		}()
	},
}

var uninstall = &cobra.Command{
	Use:               "uninstall",
	Short:             "Uninstall Parseable on kubernetes cluster",
	Long:              "\nuninstall command is used to uninstall Parseable oss/enterprise on k8s cluster.",
	PersistentPreRunE: combinedPreRun,
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		if os.Getenv("PB_ANALYTICS") == "disable" {
			return
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			analytics.PostRunAnalytics(cmd, "uninstall", args)
		}()
	},
}

func main() {
	profile.AddCommand(pb.AddProfileCmd)
	profile.AddCommand(pb.RemoveProfileCmd)
	profile.AddCommand(pb.ListProfileCmd)
	profile.AddCommand(pb.DefaultProfileCmd)

	user.AddCommand(pb.AddUserCmd)
	user.AddCommand(pb.RemoveUserCmd)
	user.AddCommand(pb.ListUserCmd)
	user.AddCommand(pb.SetUserRoleCmd)

	role.AddCommand(pb.AddRoleCmd)
	role.AddCommand(pb.RemoveRoleCmd)
	role.AddCommand(pb.ListRoleCmd)

	stream.AddCommand(pb.AddStreamCmd)
	stream.AddCommand(pb.RemoveStreamCmd)
	stream.AddCommand(pb.ListStreamCmd)
	stream.AddCommand(pb.StatStreamCmd)

	query.AddCommand(pb.QueryCmd)
	query.AddCommand(pb.SavedQueryList)

	schema.AddCommand(pb.GenerateSchemaCmd)
	schema.AddCommand(pb.CreateSchemaCmd)

	cluster.AddCommand(pb.InstallOssCmd)
	cluster.AddCommand(pb.ListOssCmd)
	cluster.AddCommand(pb.ShowValuesCmd)
	cluster.AddCommand(pb.UninstallOssCmd)

	list.AddCommand(pb.ListOssCmd)

	uninstall.AddCommand(pb.UninstallOssCmd)

	show.AddCommand(pb.ShowValuesCmd)

	cli.AddCommand(profile)
	cli.AddCommand(query)
	cli.AddCommand(stream)
	cli.AddCommand(user)
	cli.AddCommand(role)
	cli.AddCommand(pb.TailCmd)
	cli.AddCommand(cluster)

	cli.AddCommand(pb.AutocompleteCmd)

	// Set as command
	pb.VersionCmd.Run = func(_ *cobra.Command, _ []string) {
		pb.PrintVersion(Version, Commit)
	}

	cli.AddCommand(pb.VersionCmd)
	// set as flag
	cli.Flags().BoolP(versionFlag, versionFlagShort, false, "Print version")

	cli.CompletionOptions.HiddenDefaultCmd = true

	// create a default profile if file does not exist
	if previousConfig, err := config.ReadConfigFromFile(); os.IsNotExist(err) {
		conf := config.Config{
			Profiles:       map[string]config.Profile{"demo": defaultInitialProfile()},
			DefaultProfile: "demo",
		}
		err = config.WriteConfigToFile(&conf)
		if err != nil {
			fmt.Printf("failed to write to file %v\n", err)
			os.Exit(1)
		}
	} else {
		// Only update the "demo" profile without overwriting other profiles
		demoProfile, exists := previousConfig.Profiles["demo"]
		if exists {
			// Update fields in the demo profile only
			demoProfile.URL = "http://demo.parseable.com"
			demoProfile.Username = "admin"
			demoProfile.Password = "admin"
			previousConfig.Profiles["demo"] = demoProfile
		} else {
			// Add the "demo" profile if it doesn't exist
			previousConfig.Profiles["demo"] = defaultInitialProfile()
			previousConfig.DefaultProfile = "demo" // Optional: set as default if needed
		}

		// Write the updated configuration back to file
		err = config.WriteConfigToFile(previousConfig)
		if err != nil {
			fmt.Printf("failed to write to existing file %v\n", err)
			os.Exit(1)
		}
	}

	err := cli.Execute()
	if err != nil {
		os.Exit(1)
	}
	wg.Wait()
}

// Wrapper to combine existing pre-run logic and ULID check
func combinedPreRun(cmd *cobra.Command, args []string) error {
	err := pb.PreRunDefaultProfile(cmd, args)
	if err != nil {
		return fmt.Errorf("error initializing default profile: %w", err)
	}

	if err := analytics.CheckAndCreateULID(cmd, args); err != nil {
		return fmt.Errorf("error while creating ulid: %v", err)
	}

	return nil
}
