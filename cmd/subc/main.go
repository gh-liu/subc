package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	subclient "github.com/gh-liu/subc/client"
	"github.com/gh-liu/subc/protocol"
	builtintemplate "github.com/gh-liu/subc/template"
)

func main() {
	clientName := flag.String("c", "shadowrocket", "subscription client profile: shadowrocket or v2ray")
	templatePath := flag.String("t", "", "Go template file used to render parsed nodes")
	builtinTemplate := flag.String("T", "", "built-in template name used to render parsed nodes")
	flag.Parse()
	if flag.NArg() != 1 || validateTemplateFlags(*templatePath, *builtinTemplate) != nil {
		fmt.Fprintf(os.Stderr, "usage: %s [-c shadowrocket|v2ray] [-t template.gotmpl|-T %s] <subscription-url>\n", os.Args[0], strings.Join(builtintemplate.Names(), "|"))
		os.Exit(2)
	}
	client, err := subclient.Get(*clientName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	nodes, err := fetchAndParse(flag.Arg(0), client)
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

var defaultHTTPClient = &http.Client{Timeout: 30 * time.Second}

func fetchAndParse(subscriptionURL string, client subclient.Client) ([]protocol.Node, error) {
	req, err := http.NewRequest(http.MethodGet, subscriptionURL, nil)
	if err != nil {
		return nil, err
	}
	client.PrepareRequest(req)

	resp, err := defaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("fetch subscription: status %s", resp.Status)
	}
	return client.Parse(resp.Body)
}
