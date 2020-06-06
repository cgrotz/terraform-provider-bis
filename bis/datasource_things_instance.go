package bis

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func datasourceThings() *schema.Resource {
	return &schema.Resource{
		Read: resourceThingsRead,
		Schema: map[string]*schema.Schema{
			"things_solution_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"api_token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceThingsRead(d *schema.ResourceData, meta interface{}) error {
	thingsSolutionID := d.Get("things_solution_id").(string)
	d.SetId(thingsSolutionID)
	apiToken := d.Get("api_token").(string)
	d.Set("api_token", apiToken)
	return nil
}
