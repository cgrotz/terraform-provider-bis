package bis

import (
	"fmt"
	"os"

	"github.com/cgrotz/terraform-provider-bis/bis/config"
	"github.com/cgrotz/terraform-provider-bis/bis/datasources"
	"github.com/cgrotz/terraform-provider-bis/bis/resources"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider for Bosch IoT Suite
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The OAuth client's id for access the API",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
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
	clientID := d.Get("client_id").(string)
	if clientID == "" {
		clientID = os.Getenv("CLIENT_ID")
	}
	clientSecret := d.Get("client_secret").(string)
	if clientSecret == "" {
		clientSecret = os.Getenv("CLIENT_SECRET")
	}

	if clientID == "" && clientSecret == "" {
		return nil, fmt.Errorf("OAuth clientId and clientSecret must be set")
	}

	config := config.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return config, nil
}
