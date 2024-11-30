package cmd

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/spf13/cobra"
)

var day string

type MetaInfo struct {
	Package string
	Path    string
}

var rootCmd = &cobra.Command{
	Use:   "AOC",
	Short: "Generate Advent of code templates",
	Run: func(cmd *cobra.Command, _ []string) {
		isNumeric := regexp.MustCompile(`^\d+$`).MatchString(day)
		if !isNumeric {
			fmt.Println("Given day is not valid numberic value")
			os.Exit(1)
		}

		if len(day) == 1 {
			day = "0" + day
		}

		if err := initChallengeEnv(); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println("Success!")
	},
}

func initChallengeEnv() error {
	dir, _ := os.Getwd()
	dayPath := fmt.Sprintf("days/%s", day)

	path := fmt.Sprintf("%s/%s", dir, dayPath)
	if err := os.MkdirAll(path, 0777); err != nil {
		return fmt.Errorf("cannot create folders: %s", err.Error())
	}

	if _, err := os.Create(filepath.Join(path, "test")); err != nil {
		return fmt.Errorf("cannot create (test) input file: %s", err.Error())
	}
	if _, err := os.Create(filepath.Join(path, "final")); err != nil {
		return fmt.Errorf("cannot create (final) input file: %s", err.Error())
	}

	f, err := os.Create(filepath.Join(path, "main.go"))
	if err != nil {
		return fmt.Errorf("cannot create (main.go) file: %s", err.Error())
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/templates/main.tmpl", dir))
	if err != nil {
		return fmt.Errorf("could not handle template: %s", err.Error())
	}

	buf := new(bytes.Buffer)
	meta := MetaInfo{Package: day, Path: dayPath}
	if err := tmpl.Execute(buf, meta); err != nil {
		return fmt.Errorf("could not parse template with data: %s", err.Error())
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("could not write file: %s", err.Error())
	}

	return nil
}

func Execute() {
	rootCmd.Flags().StringVarP(&day, "day", "d", "", "Advent of code current day challenge")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
