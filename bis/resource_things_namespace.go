package bis

import (
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceThingsNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceThingsNamespaceCreate,
		Read:   resourceThingsNamespaceRead,
		Update: resourceThingsNamespaceUpdate,
		Delete: resourceThingsNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"things_solution_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"api_token": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceThingsNamespaceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	namespace := d.Get("namespace").(string)
	defaultNamespace := d.Get("default").(bool)
	err := CreateNamespace(config, d.Get("api_token").(string), d.Get("things_solution_id").(string), namespace, defaultNamespace)
	if err != nil {
		log.Printf("[ERROR] failed creating namespace %s", err)
		return err
	}
	d.SetId(namespace)
	return resourceThingsNamespaceRead(d, meta)
}

func resourceThingsNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	//config := meta.(Config)
	namespaceID := d.Get("namespace").(string)
	client := resty.New()
	resp, err := client.R().SetPathParams(map[string]string{
		"solutionId":  d.Get("things_solution_id").(string),
		"namespaceId": namespaceID,
	}).Get("https://things.eu-1.bosch-iot-suite.com/api/2/solutions/{solutionId}/namespaces/{namespaceId}")
	// .SetResult(things.NamespaceEntry{})
	if err != nil {
		log.Printf("[ERROR] failed reading namespace from IoT Things API %s", err)
		return err
	}
	log.Printf("[WARN] Trying to retrieve namespace %d", resp.StatusCode())
	if resp.StatusCode() == 200 {
		d.SetId(namespaceID)
		d.Set("namespace", namespaceID)
	} else {
		d.SetId("") // Empty ID means that the resource doesn't exist (anymore)
	}
	return nil
}

func resourceThingsNamespaceUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("namespace") {
		// Try updating the namespace
		if err := updateNamespace(d, m); err != nil {
			return err
		}

		d.SetPartial("namespace")
	}

	// If we were to return here, before disabling partial mode below,
	// then only the "namespace" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceThingsNamespaceRead(d, m)
}

func updateNamespace(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceThingsNamespaceDelete(d *schema.ResourceData, m interface{}) error {

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
