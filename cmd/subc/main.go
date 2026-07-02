package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/gh-liu/subc"
	builtintemplate "github.com/gh-liu/subc/template"
)

func main() {
	templatePath := flag.String("t", "", "Go template file used to render parsed nodes")
	builtinTemplate := flag.String("T", "", "built-in template name used to render parsed nodes")
	flag.Parse()
	if flag.NArg() != 1 || validateTemplateFlags(*templatePath, *builtinTemplate) != nil {
		fmt.Fprintf(os.Stderr, "usage: %s [-t template.gotmpl|-T singbox] <subscription-url>\n", os.Args[0])
		os.Exit(2)
	}
	nodes, err := subc.FetchAndParse(flag.Arg(0), subc.ShadowrocketClient{})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if *builtinTemplate != "" {
		if err := builtintemplate.RenderBuiltin(os.Stdout, *builtinTemplate, nodes); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	if *templatePath != "" {
		templateSource, err := os.ReadFile(*templatePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := builtintemplate.Render(os.Stdout, string(templateSource), nodes); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		return
	}
	out, err := json.MarshalIndent(nodes, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}

func validateTemplateFlags(templatePath, builtinTemplate string) error {
	if templatePath != "" && builtinTemplate != "" {
		return fmt.Errorf("-t and -T cannot be used together")
	}
	return nil
}
