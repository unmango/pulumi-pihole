package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ryanwholey/terraform-provider-pihole/internal/provider"
)

func NewProvider() *schema.Provider {
	p, _ := provider.New()
	return p
}
