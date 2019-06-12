package routeros

import (
	"fmt"
	"log"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDNSStaticRecord() *schema.Resource {
	return &schema.Resource{
		Create: resourceDNSStaticRecordCreate,
		Read:   resourceDNSStaticRecordRead,
		Update: resourceDNSStaticRecordUpdate,
		Delete: resourceDNSStaticRecordDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "managed_by_terraform",
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "30m",
			},
		},
	}
}

func resourceDNSStaticRecordCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDNSStaticRecordCreate")

	r := &rc.ResourceDNSStaticRecord{
		Address:  d.Get("address").(string),
		Disabled: d.Get("disabled").(bool),
		Comment:  d.Get("comment").(string),
		Name:     d.Get("name").(string),
		TTL:      d.Get("ttl").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceDNSStaticRecordRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDNSStaticRecordRead")

	r := &rc.ResourceDNSStaticRecord{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceDNSStaticRecord); ok {
		d.Set("address", res.Address)
		d.Set("disabled", res.Disabled)
		d.Set("comment", res.Comment)
		d.Set("name", res.Name)
		d.Set("ttl", res.TTL)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceDNSStaticRecordUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDNSStaticRecordUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceDNSStaticRecord{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	n := &rc.ResourceDNSStaticRecord{
		ID:       d.Id(),
		Address:  d.Get("address").(string),
		Disabled: d.Get("disabled").(bool),
		Comment:  d.Get("comment").(string),
		Name:     d.Get("name").(string),
		TTL:      d.Get("ttl").(string),
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

func resourceDNSStaticRecordDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDNSStaticRecordDelete")

	r := &rc.ResourceDNSStaticRecord{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
