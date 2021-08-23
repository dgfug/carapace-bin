package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	names, descriptions := readCompleters()

	imports := make([]string, 0, len(names))
	for _, name := range names {
		imports = append(imports, fmt.Sprintf(`	%v "github.com/rsteube/carapace-bin/completers/%v_completer/cmd"`, varName(name), name))
	}

	formattedNames := make([]string, 0)
	for _, name := range names {
		formattedNames = append(formattedNames, fmt.Sprintf("\t\"%v\",", name))
	}

	maxlen := 0
	for _, name := range names {
		if l := len(name); l > maxlen {
			maxlen = l
		}
	}

	formattedDescriptions := make([]string, 0)
	for _, name := range names {
		formattedDescriptions = append(formattedDescriptions, fmt.Sprintf(`	%--`+strconv.Itoa(maxlen+4)+`v"%v",`, fmt.Sprintf(`"%v": `, name), descriptions[name]))
	}

	cases := make([]string, 0)
	for _, name := range names {
		cases = append(cases,
			fmt.Sprintf(`	case "%v":`, name),
			fmt.Sprintf(`		%v.Execute()`, varName(name)),
		)
	}

	content := fmt.Sprintf(`package cmd

import (
%v)

var completers = []string{
%v}

var descriptions = map[string]string{
%v}

func executeCompleter(completer string) {
	switch completer {
%v	}
}
`,
		fmt.Sprintln(strings.Join(imports, "\n")),
		fmt.Sprintln(strings.Join(formattedNames, "\n")),
		fmt.Sprintln(strings.Join(formattedDescriptions, "\n")),
		fmt.Sprintln(strings.Join(cases, "\n")),
	)
	if root, err := rootDir(); err == nil {
		ioutil.WriteFile(root+"/cmd/carapace/cmd/completers.go", []byte("// +build !release\n\n"+content), 0644)
		ioutil.WriteFile(root+"/cmd/carapace/cmd/completers_release.go", []byte("// +build release\n\n"+strings.Replace(content, "/completers/", "/completers_release/", -1)), 0644)
		os.RemoveAll(root + "/completers_release")
	}
}

func varName(name string) string {
	if name == "go" {
		return "_go"
	}
	return strings.Replace(name, "-", "_", -1)
}

func readCompleters() ([]string, map[string]string) {
	names := make([]string, 0)
	descriptions := make(map[string]string)
	if root, err := rootDir(); err == nil {
		if files, err := ioutil.ReadDir(root + "/completers/"); err == nil {
			for _, file := range files {
				if file.IsDir() && strings.HasSuffix(file.Name(), "_completer") {
					name := strings.TrimSuffix(file.Name(), "_completer")
					description := readDescription(root, file.Name())
					names = append(names, name)
					descriptions[name] = description
				}
			}
		}
	}
	return names, descriptions
}

func readDescription(root string, completer string) string {
	if content, err := ioutil.ReadFile(fmt.Sprintf("%v/completers/%v/cmd/root.go", root, completer)); err == nil {
		re := regexp.MustCompile("^\tShort: +\"(?P<description>.*)\",$")
		for _, line := range strings.Split(string(content), "\n") {
			if re.MatchString(line) {
				return re.FindStringSubmatch(line)[1]
			}
		}
	}
	return ""
}

func rootDir() (string, error) {
	if output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		return "", err
	} else {
		return strings.Split(string(output), "\n")[0], nil
	}
}
