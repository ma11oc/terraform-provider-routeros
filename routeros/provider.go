package routeros

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"

	rc "github.com/ma11oc/go-routerosclient"
)

var RouterOSClient *rc.Client

// Provider ...
// FIXME: add provider description
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_DEFAULT_ADDR", nil),
				Description: "RouterOS address and port",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_DEFAULT_USERNAME", nil),
				Description: "User name",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password",
				DefaultFunc: schema.EnvDefaultFunc("ROS_DEFAULT_PASSWORD", nil),
			},
			"async": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("ROS_DEFAULT_ASYNC", false),
				Description: "Use async code",
			},
			/*
			 * "usetls": {
			 *     Type:        schema.TypeBool,
			 *     Optional:    true,
			 *     DefaultFunc: schema.EnvDefaultFunc("ROS_DEFAULT_USE_TLS", false),
			 *     Description: "Use TLS",
			 * },
			 */
		},

		ResourcesMap: map[string]*schema.Resource{
			"ros_dns_static_record":      resourceDNSStaticRecord(),
			"ros_dhcp_server":            resourceDHCPServer(),
			"ros_dhcp_server_network":    resourceDHCPServerNetwork(),
			"ros_dhcp_server_lease":      resourceDHCPServerLease(),
			"ros_dhcp_server_option":     resourceDHCPServerOption(),
			"ros_dhcp_server_option_set": resourceDHCPServerOptionSet(),
			"ros_interface_bridge":       resourceInterfaceBridge(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		address:  d.Get("address").(string),
		username: d.Get("username").(string),
		password: d.Get("password").(string),
		async:    d.Get("async").(bool),
	}

	return config.Client()
}
