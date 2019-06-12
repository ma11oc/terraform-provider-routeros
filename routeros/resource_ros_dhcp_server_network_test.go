package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceDHCPServerNetwork(t *testing.T) {
	var r = &rc.ResourceDHCPServerNetwork{
		Address: "192.168.0.0/24",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDHCPServerNetworkDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprint(`
					resource "ros_dhcp_server_network" "tf-test-dhcp-network" {
						address = "192.168.0.0/24"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerNetworkExists("ros_dhcp_server_network.tf-test-dhcp-network", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "address", "192.168.0.0/24"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_dhcp_server_option" "tf-test-next-server" {
						code = 66
						name = "next-server"
						value = "'192.168.0.254'"
					}

					resource "ros_dhcp_server_network" "tf-test-dhcp-network" {
						address = "192.168.0.0/24"
						boot_file_name = "tftpboot/pxelinux.0"
						dhcp_option = "${ros_dhcp_server_option.tf-test-next-server.name}"
						domain = "example.com"
						dns_server = "192.168.0.1"
						gateway = "192.168.0.1"
						netmask = "24"
						next_server = "192.168.0.10"
						ntp_server = "192.168.0.10"
						wins_server = "192.168.0.10"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerNetworkExists("ros_dhcp_server_network.tf-test-dhcp-network", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "address", "192.168.0.0/24"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "boot_file_name", "tftpboot/pxelinux.0"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "dhcp_option", "next-server"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "domain", "example.com"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "dns_server", "192.168.0.1"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "gateway", "192.168.0.1"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "netmask", "24"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "next_server", "192.168.0.10"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "ntp_server", "192.168.0.10"),
					resource.TestCheckResourceAttr("ros_dhcp_server_network.tf-test-dhcp-network", "wins_server", "192.168.0.10"),
				),
			},
		},
	})
}

func resourceDHCPServerNetworkExists(n string, r *rc.ResourceDHCPServerNetwork) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dhcp option ID is set")
		}

		client := testAccProvider.Meta().(*rc.Client)

		if err, ok := client.CheckResourceExists(r); !ok {
			return err
		}

		return nil
	}
}

func testAccCheckDHCPServerNetworkDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_dhcp_server_network" {
			continue
		}

		r := &rc.ResourceDHCPServerNetwork{
			ID: rs.Primary.ID,
		}

		if _, ok := client.CheckResourceExists(r); ok {
			return fmt.Errorf(
				"Error waiting for option with id %v to be destroyed",
				rs.Primary.ID,
			)
		}

	}

	return nil
}
