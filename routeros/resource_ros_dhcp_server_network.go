package routeros

import (
	"fmt"
	"log"

	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceDHCPServerNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceDHCPServerNetworkCreate,
		Read:   resourceDHCPServerNetworkRead,
		Update: resourceDHCPServerNetworkUpdate,
		Delete: resourceDHCPServerNetworkDelete,

		Schema: map[string]*schema.Schema{
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  false,
			},
			"boot_file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"comment": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "managed_by_terraform",
			},
			"dhcp_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhcp_option_set": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"netmask": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"next_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ntp_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wins_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDHCPServerNetworkCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerNetworkCreate")

	r := &rc.ResourceDHCPServerNetwork{
		Address:       d.Get("address").(string),
		BootFileName:  d.Get("boot_file_name").(string),
		Comment:       d.Get("comment").(string),
		DHCPOption:    d.Get("dhcp_option").(string),
		DHCPOptionSet: d.Get("dhcp_option_set").(string),
		Domain:        d.Get("domain").(string),
		DNSServer:     d.Get("dns_server").(string),
		Gateway:       d.Get("gateway").(string),
		Netmask:       d.Get("netmask").(string),
		NextServer:    d.Get("next_server").(string),
		NTPServer:     d.Get("ntp_server").(string),
		WINSServer:    d.Get("wins_server").(string),
	}

	client := m.(*rc.Client)

	id, err := client.CreateResource(r)
	if err != nil {
		return err
	}

	d.SetId(id)

	return nil
}

func resourceDHCPServerNetworkRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerNetworkRead")

	r := &rc.ResourceDHCPServerNetwork{
		ID: d.Id(),
	}

	client := m.(*rc.Client)

	obj, err := client.ReadResource(r)
	if err != nil {
		return err
	}

	if res, ok := obj.(*rc.ResourceDHCPServerNetwork); ok {
		d.Set("address", res.Address)
		d.Set("boot_file_name", res.BootFileName)
		d.Set("comment", res.Comment)
		d.Set("dhcp_option", res.DHCPOption)
		d.Set("dhcp_option_set", res.DHCPOptionSet)
		d.Set("domain", res.Domain)
		d.Set("dns_server", res.DNSServer)
		d.Set("gateway", res.Gateway)
		d.Set("netmask", res.Netmask)
		d.Set("next_server", res.NextServer)
		d.Set("ntp_server", res.NTPServer)
		d.Set("wins_server", res.WINSServer)
	} else {
		return fmt.Errorf("unable to cast resource")
	}

	return nil
}

func resourceDHCPServerNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerNetworkUpdate")
	// Partial state is not required since it's an atomic call

	o := &rc.ResourceDHCPServerNetwork{
		ID: d.Id(),
	}

	n := &rc.ResourceDHCPServerNetwork{
		ID:            d.Id(),
		Address:       d.Get("address").(string),
		BootFileName:  d.Get("boot_file_name").(string),
		Comment:       d.Get("comment").(string),
		DHCPOption:    d.Get("dhcp_option").(string),
		DHCPOptionSet: d.Get("dhcp_option_set").(string),
		Domain:        d.Get("domain").(string),
		DNSServer:     d.Get("dns_server").(string),
		Gateway:       d.Get("gateway").(string),
		Netmask:       d.Get("netmask").(string),
		NextServer:    d.Get("next_server").(string),
		NTPServer:     d.Get("ntp_server").(string),
		WINSServer:    d.Get("wins_server").(string),
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

func resourceDHCPServerNetworkDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[DEBUG] resourceDHCPServerNetworkDelete")

	r := &rc.ResourceDHCPServerNetwork{
		ID: d.Id(),
	}

	client := m.(*rc.Client)

	if err, ok := client.DeleteResource(r); !ok {
		return err
	}

	return nil
}
