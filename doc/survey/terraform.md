# Terraform

## Plugin

- it use binary for plugin, but not using https://golang.org/pkg/plugin/
- it use rpc https://github.com/hashicorp/go-plugin, when the binary starts, it connects to a server
- https://www.terraform.io/docs/plugins/basics.html#developing-a-plugin
- https://github.com/terraform-providers/terraform-provider-aws