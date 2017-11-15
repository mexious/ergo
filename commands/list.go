package commands

import (
	"fmt"

	"github.com/cristianoliveira/ergo/proxy"
)

// ListCommand lists all configured apps local url and its original urls.
// Usage:
// `ergo list`
type ListCommand struct{}

// Execute apply the ListCommand
func (c ListCommand) Execute(config *proxy.Config) (string, error) {
	output := "Ergo Proxy current list:\n"

	for name, service := range config.Services {
		localURL := `http://` + name + config.Domain

		output = fmt.Sprintf("%s - %s -> %s \n", output, localURL, service.URL)
	}

	return output, nil
}
