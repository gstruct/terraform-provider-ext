package main

import (
    "github.com/hashicorp/terraform/plugin"
    "github.com/gstruct/terraform-provider-ext/ext"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: ext.Provider,
    })
}
