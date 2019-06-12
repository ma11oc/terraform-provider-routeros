package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceDHCPServerLease(t *testing.T) {
	var r = &rc.ResourceDHCPServerLease{
		Address:    "169.254.169.254",
		Server:     "tf-test-dhcp-server",
		MacAddress: "00:11:22:33:44:55",
		Disabled:   false,
		ClientID:   "tf-test-dhcp-server-lease",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDHCPServerLeaseDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprint(`
					resource "ros_interface_bridge" "tf-test-br" {
						name = "tf-test-br"
						fast_forward = true
						disabled = false
					}

					resource "ros_dhcp_server" "tf-test-dhcp-server" {
						name = "tf-test-dhcp-server"
						interface = "${ros_interface_bridge.tf-test-br.name}"
					}

					resource "ros_dhcp_server_lease" "tf-test-dhcp-server-lease" {
						address = "169.254.169.254"
						server = "${ros_dhcp_server.tf-test-dhcp-server.name}"
						mac = "00:11:22:33:44:55"
						disabled = false
						client_id = "tf-test-dhcp-server-lease"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerLeaseExists("ros_dhcp_server_lease.tf-test-dhcp-server-lease", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_lease.tf-test-dhcp-server-lease", "disabled", "false"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_interface_bridge" "tf-test-br" {
						name = "tf-test-br"
						fast_forward = true
						disabled = false
					}

					resource "ros_dhcp_server" "tf-test-dhcp-server" {
						name = "tf-test-dhcp-server"
						interface = "${ros_interface_bridge.tf-test-br.name}"
					}

					resource "ros_dhcp_server_lease" "tf-test-dhcp-server-lease" {
						address = "169.254.169.254"
						server = "${ros_dhcp_server.tf-test-dhcp-server.name}"
						mac = "00:11:22:33:44:55"
						disabled = true
						client_id = "tf-test-dhcp-server-lease"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerLeaseExists("ros_dhcp_server_lease.tf-test-dhcp-server-lease", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_lease.tf-test-dhcp-server-lease", "disabled", "true"),
				),
			},
		},
	})
}

func resourceDHCPServerLeaseExists(n string, r *rc.ResourceDHCPServerLease) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No dhcp lease ID is set")
		}

		client := testAccProvider.Meta().(*rc.Client)

		if err, ok := client.CheckResourceExists(r); !ok {
			return err
		}

		return nil
	}
}

func testAccCheckDHCPServerLeaseDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_dhcp_server_lease" {
			continue
		}

		r := &rc.ResourceDHCPServerLease{
			ID: rs.Primary.ID,
		}

		if _, ok := client.CheckResourceExists(r); ok {
			return fmt.Errorf(
				"Error waiting for lease with id %v to be destroyed",
				rs.Primary.ID,
			)
		}

	}

	return nil
}
