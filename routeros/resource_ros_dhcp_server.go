package routeros

import (
	"fmt"
	"log"

	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDHCPServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPServerCreate,
		Read:   resourceDHCPServerRead,
		Update: resourceDHCPServerUpdate,
		Delete: resourceDHCPServerDelete,

		Schema: map[string]*schema.Schema{
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDHCPServerCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerCreate")

	r := &rc.ResourceDHCPServer{
		Disabled:  d.Get("disabled").(bool),
		Interface: d.Get("interface").(string),
		Name:      d.Get("name").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceDHCPServerRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerRead")

	r := &rc.ResourceDHCPServer{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceDHCPServer); ok {
		d.Set("disabled", res.Disabled)
		d.Set("interface", res.Interface)
		d.Set("name", res.Name)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceDHCPServerUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceDHCPServer{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	n := &rc.ResourceDHCPServer{
		ID:        d.Id(),
		Disabled:  d.Get("disabled").(bool),
		Interface: d.Get("interface").(string),
		Name:      d.Get("name").(string),
	}

	/*
	 *     if d.HasChange("disabled") {
	 *         n.Disabled = d.Get("disabled").(bool)
	 *     }
	 *
	 *     if d.HasChange("interface") {
	 *         n.Comment = d.Get("interface").(string)
	 *     }
	 *
	 *     if d.HasChange("name") {
	 *         n.Comment = d.Get("name").(string)
	 *     }
	 */

	client := m.(*rc.Client)

	if err, ok := client.UpdateResource(o, n); !ok {
		return err
	}

	return nil
}

func resourceDHCPServerDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerDelete")

	r := &rc.ResourceDHCPServer{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
