package datasources

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// DatasourceThings creates the schema description for the things resource
func DatasourceThings() *schema.Resource {
	return &schema.Resource{
		Read: resourceThingsRead,
		Schema: map[string]*schema.Schema{
			"solution_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_token": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func resourceThingsRead(d *schema.ResourceData, meta interface{}) error {
	solutionID := d.Get("solution_id").(string)
	d.SetId(solutionID)
	apiToken := d.Get("api_token").(string)
	d.Set("api_token", apiToken)
	return nil
}
