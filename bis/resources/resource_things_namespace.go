package resources

import (
	"log"

	"github.com/cgrotz/terraform-provider-bis/bis/client"
	"github.com/cgrotz/terraform-provider-bis/bis/config"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// ResourceThingsNamespace creates the schema description for the things namespace resource
func ResourceThingsNamespace() *schema.Resource {
	return &schema.Resource{
		Create: resourceThingsNamespaceCreate,
		Read:   resourceThingsNamespaceRead,
		Update: resourceThingsNamespaceUpdate,
		Delete: resourceThingsNamespaceDelete,

		Schema: map[string]*schema.Schema{
			"solution_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
	config := meta.(config.Config)
	namespace := d.Get("namespace").(string)
	defaultNamespace := d.Get("default").(bool)
	err := client.CreateNamespace(config, d.Get("solution_id").(string), namespace, defaultNamespace)
	if err != nil {
		log.Printf("[ERROR] failed creating namespace %s", err)
		return err
	}
	d.SetId(namespace)
	d.Set("default", defaultNamespace)
	return resourceThingsNamespaceRead(d, meta)
}

func resourceThingsNamespaceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config.Config)
	namespace := d.Get("namespace").(string)
	namespaceEntry, err := client.RetrieveNamespace(config, d.Get("solution_id").(string), namespace)
	if err != nil {
		log.Printf("[ERROR] failed creating namespace %s", err)
		return err
	}
	d.SetId(namespace)
	d.Set("default", namespaceEntry.Default)
	return nil
}

func resourceThingsNamespaceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config.Config)
	// Enable partial state mode
	d.Partial(true)

	if d.HasChange("default") {
		namespace := d.Get("namespace").(string)
		defaultNamespace := d.Get("default").(bool)
		err := client.CreateNamespace(config, d.Get("solution_id").(string), namespace, defaultNamespace)
		if err != nil {
			log.Printf("[ERROR] failed creating namespace %s", err)
			return err
		}
		d.Set("default", defaultNamespace)

		d.SetPartial("default")
	}

	// If we were to return here, before disabling partial mode below,
	// then only the "namespace" field would be saved.

	// We succeeded, disable partial mode. This causes Terraform to save
	// all fields again.
	d.Partial(false)

	return resourceThingsNamespaceRead(d, meta)
}

func resourceThingsNamespaceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(config.Config)
	namespace := d.Get("namespace").(string)
	err := client.DeleteNamespace(config, d.Get("solution_id").(string), namespace)
	if err != nil {
		log.Printf("[ERROR] failed creating namespace %s", err)
		return err
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")
	return nil
}
