package routeros

import (
	"fmt"
	"log"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDHCPServerOption() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPServerOptionCreate,
		Read:   resourceDHCPServerOptionRead,
		Update: resourceDHCPServerOptionUpdate,
		Delete: resourceDHCPServerOptionDelete,

		Schema: map[string]*schema.Schema{
			"code": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDHCPServerOptionCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionCreate")

	r := &rc.ResourceDHCPServerOption{
		Code:  d.Get("code").(int),
		Name:  d.Get("name").(string),
		Value: d.Get("value").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceDHCPServerOptionRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionRead")

	r := &rc.ResourceDHCPServerOption{
		ID: d.Id(),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceDHCPServerOption); ok {
		d.Set("code", res.Code)
		d.Set("name", res.Name)
		d.Set("value", res.Value)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceDHCPServerOptionUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceDHCPServerOption{
		ID: d.Id(),
	}

	n := &rc.ResourceDHCPServerOption{
		ID:    d.Id(),
		Code:  d.Get("code").(int),
		Name:  d.Get("name").(string),
		Value: d.Get("value").(string),
	}

	client := m.(*rc.Client)

	if err, ok := client.UpdateResource(o, n); !ok {
		return err
	}

	return nil
}

func resourceDHCPServerOptionDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerOptionDelete")

	r := &rc.ResourceDHCPServerOption{
		ID: d.Id(),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
