package main

import (
	"fmt"
	"os"
	"plugin"

	"github.com/tonyhhyip/seau/api"
)

func main() {
	files := os.Args[1:]
	for _, file := range files {
		validate(file)
	}
}

func validate(name string) {
	fmt.Printf("Checking %s\n", name)
	p, err := plugin.Open(name)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	sym, err := p.Lookup(api.PluginSymbol)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	apiPlugin, ok := sym.(api.Plugin)
	if !ok {
		fmt.Println("Error: Fail to load as plugin")
		return
	}

	fmt.Printf("Validate plugin: %s(%s)\n", apiPlugin.Name(), apiPlugin.ID())
}
