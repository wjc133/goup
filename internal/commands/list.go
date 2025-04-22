package commands

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List all installed Go",
		Long:    "List all installed Go versions.",
		RunE:    runList,
	}
}

func runList(cmd *cobra.Command, args []string) error {
	vers, err := listGoVers()
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Version", "Active"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.SetAlignment(tablewriter.ALIGN_CENTER)

	for _, ver := range vers {
		if ver.Current {
			table.Append([]string{ver.Ver, "*"})
		} else {
			table.Append([]string{ver.Ver, ""})
		}
	}

	table.Render()

	return nil
}

type goVer struct {
	Ver     string
	Current bool
}

func listGoVers() ([]goVer, error) {
	goVersions := make([]goVer, 0)
	baseDir := GoBaseDir()
	baseDirVersions, err := findGoVers(baseDir)
	if err != nil {
		return nil, err
	}
	goVersions = append(goVersions, baseDirVersions...)
	return goVersions, nil
}

func findGoVers(dirPath string) ([]goVer, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	current, err := currentGoVersion()
	if err != nil {
		return nil, err
	}

	var vers []goVer
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if strings.HasPrefix(file.Name(), "go") || file.Name() == "gotips" {
			vers = append(vers, goVer{
				Ver:     strings.TrimPrefix(file.Name(), "go"),
				Current: current == file.Name(),
			})
		}
	}

	return vers, nil
}

func currentGoVersion() (string, error) {
	current := GoupCurrentDir()
	goroot, err := os.Readlink(current)
	if err != nil {
		return "", err
	}

	return filepath.Base(goroot), nil
}
