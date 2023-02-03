// nolint:deadcode
package utils

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	r "runtime"

	"github.com/bndr/gotabulate"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/manifoldco/promptui"
)

// Needle use for switch
type Needle struct {
	Name string
	// Cluster string
	// User    string
	// Center  string
	Desc string
}

// Namespaces namespaces struct
type Namespaces struct {
	Name    string
	Default bool
}

// Copied from https://github.com/kubernetes/kubernetes
// /blob/master/pkg/kubectl/util/hash/hash.go
func hEncode(hex string) (string, error) {
	if len(hex) < 10 {
		return "", fmt.Errorf(
			"input length must be at least 10")
	}
	enc := []rune(hex[:10])
	for i := range enc {
		switch enc[i] {
		case '0':
			enc[i] = 'g'
		case '1':
			enc[i] = 'h'
		case '3':
			enc[i] = 'k'
		case 'a':
			enc[i] = 'm'
		case 'e':
			enc[i] = 't'
		}
	}
	return string(enc), nil
}

// Hash returns the hex form of the sha256 of the argument.
func Hash(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

// HashSufString return the string of HashSuf.
func HashSufString(data string) string {
	sum, _ := hEncode(Hash(data))
	return sum
}

// PrintTable generate table
func PrintTable(table interface{}, headers []string) error {
	tabulate := gotabulate.Create(table)
	tabulate.SetHeaders(headers)
	// Turn On String Wrapping
	tabulate.SetWrapStrings(true)
	// Render the table
	tabulate.SetAlign("center")
	// Set Max Cell Size
	tabulate.SetMaxCellSize(60)
	fmt.Println(tabulate.Render("grid", "left"))
	return nil
}

// SelectUI output select ui
func SelectUI(kubeItems []Needle, label string, uiSize int) int {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U0001F63C {{ .Name | red }}{{ .Center | red}}",
		Inactive: "  {{ .Name | cyan }}{{ .Center | red}}",
		Selected: "\U0001F638 Select:{{ .Name | green }}",
		Details: `
--------- Info ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Cluster:" | faint }}	{{ .Cluster }}
{{ "User:" | faint }}	{{ .User }}`,
	}
	searcher := func(input string, index int) bool {
		pepper := kubeItems[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)
		if input == "q" && name == "<exit>" {
			return true
		}
		return strings.Contains(name, input)
	}
	prompt := promptui.Select{
		Label:     label,
		Items:     kubeItems,
		Templates: templates,
		Size:      uiSize,
		Searcher:  searcher,
	}
	i, _, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt fail %v\n", err)
	}
	if kubeItems[i].Name == "<Exit>" {
		fmt.Println("Exited.")
		os.Exit(1)
	}
	return i
}

// PromptUI output prompt ui
func PromptUI(label string, name string) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New("context name must have more than 1 characters")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		Default:  name,
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt fail %v\n", err)
	}
	return result
}

// BoolUI output bool ui
func BoolUI(label string, uiSize int) string {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}",
		Active:   "\U0001F37A {{ . | red }}",
		Inactive: "  {{ . | cyan }}",
		Selected: "\U0001F47B {{ . | green }}",
	}
	prompt := promptui.Select{
		Label:     label,
		Items:     []string{"False", "True"},
		Templates: templates,
		Size:      uiSize,
	}
	_, obj, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt fail %v\n", err)
	}
	return obj
}

// ExitOption exit option of SelectUI
func ExitOption(kubeItems []Needle) []Needle {
	kubeItems = append(kubeItems, Needle{Name: "<Exit>", Desc: "exit the githubcm"})
	return kubeItems
}

func printService(out io.Writer, name, link string) {
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Fprint(out, name)
	ct.ResetColor()
	fmt.Fprint(out, " is running at ")
	ct.ChangeColor(ct.Yellow, false, ct.None, false)
	fmt.Fprint(out, link)
	ct.ResetColor()
	fmt.Fprintln(out, "")
}

func PrintKV(out io.Writer, prefix string, kv map[string]interface{}) {
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Fprint(out, prefix)
	ct.ResetColor()
	for k, v := range kv {
		ct.ChangeColor(ct.Cyan, false, ct.None, false)
		fmt.Fprint(out, k)
		fmt.Fprint(out, ": ")
		ct.ResetColor()
		ct.ChangeColor(ct.Yellow, false, ct.None, false)
		fmt.Fprint(out, v)
		ct.ResetColor()
		fmt.Fprint(out, " ")
	}
	fmt.Fprint(out, "\n")
}

func getFileName(path string) string {
	n := strings.Split(path, "/")
	result := strings.Split(n[len(n)-1], ".")
	return result[0]
}

// isMacOs check if current system is macOS
func isMacOs() bool {
	return r.GOOS == "darwin"
}

func PrintString(out io.Writer, name string) {
	ct.ChangeColor(ct.Green, false, ct.None, false)
	fmt.Fprint(out, name)
	ct.ResetColor()
}

func PrintYellow(out io.Writer, content string) {
	ct.ChangeColor(ct.Yellow, false, ct.None, false)
	fmt.Fprint(out, content)
	ct.ResetColor()
}

func PrintWarning(out io.Writer, name string) {
	ct.ChangeColor(ct.Red, false, ct.None, false)
	fmt.Fprint(out, name)
	ct.ResetColor()
}
