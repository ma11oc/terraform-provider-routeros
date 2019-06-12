package routeros

import (
	"fmt"
	"log"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDHCPServerLease() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPServerLeaseCreate,
		Read:   resourceDHCPServerLeaseRead,
		Update: resourceDHCPServerLeaseUpdate,
		Delete: resourceDHCPServerLeaseDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address_lists": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "managed_by_terraform",
			},
			"disabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dhcp_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"dhcp_option_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceDHCPServerLeaseCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerLeaseCreate")

	r := &rc.ResourceDHCPServerLease{
		Address:       d.Get("address").(string),
		AddressLists:  d.Get("address_lists").(string),
		ClientID:      d.Get("client_id").(string),
		Comment:       d.Get("comment").(string),
		Disabled:      d.Get("disabled").(bool),
		DHCPOption:    d.Get("dhcp_option").(string),
		DHCPOptionSet: d.Get("dhcp_option_set").(string),
		MacAddress:    d.Get("mac").(string),
		Server:        d.Get("server").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceDHCPServerLeaseRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerLeaseRead")

	r := &rc.ResourceDHCPServerLease{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceDHCPServerLease); ok {
		d.Set("address", res.Address)
		d.Set("address_lists", res.AddressLists)
		d.Set("client_id", res.ClientID)
		d.Set("comment", res.Comment)
		d.Set("disabled", res.Disabled)
		d.Set("dhcp_option", res.DHCPOption)
		d.Set("dhcp_option_set", res.DHCPOptionSet)
		d.Set("mac", res.MacAddress)
		d.Set("server", res.Server)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceDHCPServerLeaseUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerLeaseUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceDHCPServerLease{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	n := &rc.ResourceDHCPServerLease{
		ID:            d.Id(),
		Address:       d.Get("address").(string),
		AddressLists:  d.Get("address_lists").(string),
		ClientID:      d.Get("client_id").(string),
		Comment:       d.Get("comment").(string),
		Disabled:      d.Get("disabled").(bool),
		DHCPOption:    d.Get("dhcp_option").(string),
		DHCPOptionSet: d.Get("dhcp_option_set").(string),
		MacAddress:    d.Get("mac").(string),
		Server:        d.Get("server").(string),
	}

	client := m.(*rc.Client)

	if err, ok := client.UpdateResource(o, n); !ok {
		return err
	}

	return nil
}

func resourceDHCPServerLeaseDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerLeaseDelete")

	r := &rc.ResourceDHCPServerLease{
		ID:       d.Id(),
		Disabled: d.Get("disabled").(bool),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
