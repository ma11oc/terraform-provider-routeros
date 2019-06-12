package routeros

import (
	"fmt"
	"log"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInterfaceBridge() *schema.Resource {
	return &schema.Resource{
		Create: resourceInterfaceBridgeCreate,
		Read:   resourceInterfaceBridgeRead,
		Update: resourceInterfaceBridgeUpdate,
		Delete: resourceInterfaceBridgeDelete,

		Schema: map[string]*schema.Schema{
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
			"fast_forward": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"forward_delay": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "15s",
			},
			"mtu": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceInterfaceBridgeCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceInterfaceBridgeCreate")

	r := &rc.ResourceInterfaceBridge{
		Comment:      d.Get("comment").(string),
		Disabled:     d.Get("disabled").(bool),
		FastForward:  d.Get("fast_forward").(bool),
		ForwardDelay: d.Get("forward_delay").(string),
		MTU:          d.Get("mtu").(int),
		Name:         d.Get("name").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceInterfaceBridgeRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceInterfaceBridgeRead")

	r := &rc.ResourceInterfaceBridge{
		ID:          d.Id(),
		Disabled:    d.Get("disabled").(bool),
		FastForward: d.Get("fast_forward").(bool),
		MTU:         d.Get("mtu").(int),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceInterfaceBridge); ok {
		d.Set("comment", res.Comment)
		d.Set("disabled", res.Disabled)
		d.Set("fast_forward", res.FastForward)
		d.Set("forward_delay", res.ForwardDelay)
		d.Set("mtu", res.MTU)
		d.Set("name", res.Name)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceInterfaceBridgeUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceInterfaceBridgeUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceInterfaceBridge{
		ID: d.Id(),
	}

	n := &rc.ResourceInterfaceBridge{
		ID:           d.Id(),
		Comment:      d.Get("comment").(string),
		Disabled:     d.Get("disabled").(bool),
		FastForward:  d.Get("fast_forward").(bool),
		ForwardDelay: d.Get("forward_delay").(string),
		MTU:          d.Get("mtu").(int),
		Name:         d.Get("name").(string),
	}

	/*
	 *     if d.HasChange("comment") {
	 *         n.Comment = d.Get("comment").(string)
	 *     }
	 *
	 *     if d.HasChange("disabled") {
	 *         n.Disabled = d.Get("disabled").(bool)
	 *     }
	 *
	 *     if d.HasChange("fast_forward") {
	 *         n.FastForward = d.Get("fast_forward").(bool)
	 *     }
	 *
	 *     if d.HasChange("forward_delay") {
	 *         n.ForwardDelay = d.Get("forward_delay").(string)
	 *     }
	 *
	 *     if d.HasChange("mtu") {
	 *         n.MTU = d.Get("mtu").(int)
	 *     }
	 *
	 *     if d.HasChange("name") {
	 *         n.Name = d.Get("name").(string)
	 *     }
	 */

	client := m.(*rc.Client)

	if err, ok := client.UpdateResource(o, n); !ok {
		return err
	}

	return nil
}

func resourceInterfaceBridgeDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceInterfaceBridgeDelete")

	r := &rc.ResourceInterfaceBridge{
		ID:          d.Id(),
		Disabled:    d.Get("disabled").(bool),
		FastForward: d.Get("fast_forward").(bool),
		MTU:         d.Get("mtu").(int),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
