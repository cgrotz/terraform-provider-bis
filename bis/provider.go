package bis

import (
	"github.com/cgrotz/terraform-provider-bis/bis/config"
	"github.com/cgrotz/terraform-provider-bis/bis/datasources"
	"github.com/cgrotz/terraform-provider-bis/bis/resources"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider for Bosch IoT Suite
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The OAuth client's id for access the API",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The OAuth client's secret for access the API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bis_things_namespace": resources.ResourceThingsNamespace(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bis_things": datasources.DatasourceThings(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := config.Config{
		ClientID:     d.Get("client_id").(string),
		ClientSecret: d.Get("client_secret").(string),
	}

	return config, nil
}
