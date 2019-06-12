package routeros

import (
	"fmt"
	"log"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDHCPServerOptionSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPServerOptionSetCreate,
		Read:   resourceDHCPServerOptionSetRead,
		Update: resourceDHCPServerOptionSetUpdate,
		Delete: resourceDHCPServerOptionSetDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDHCPServerOptionSetCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionSetCreate")

	r := &rc.ResourceDHCPServerOptionSet{
		Name:    d.Get("name").(string),
		Options: d.Get("options").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceDHCPServerOptionSetRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionSetRead")

	r := &rc.ResourceDHCPServerOptionSet{
		ID: d.Id(),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceDHCPServerOptionSet); ok {
		d.Set("name", res.Name)
		d.Set("options", res.Options)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceDHCPServerOptionSetUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionSetUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceDHCPServerOptionSet{
		ID: d.Id(),
	}

	n := &rc.ResourceDHCPServerOptionSet{
		ID:      d.Id(),
		Name:    d.Get("name").(string),
		Options: d.Get("options").(string),
	}

	client := m.(*rc.Client)

	if err, ok := client.UpdateResource(o, n); !ok {
		return err
	}

	return nil
}

func resourceDHCPServerOptionSetDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionSetDelete")

	r := &rc.ResourceDHCPServerOptionSet{
		ID: d.Id(),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
