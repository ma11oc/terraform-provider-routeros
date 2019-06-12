package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceDHCPServerOptionSet(t *testing.T) {
	var r = &rc.ResourceDHCPServerOptionSet{
		Name:    "PXEClient",
		Options: "next-server",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDHCPServerOptionSetDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprint(`
					resource "ros_dhcp_server_option" "tf-test-next-server" {
						code = 66
						name = "next-server"
						value = "'192.168.0.254'"
					}

					resource "ros_dhcp_server_option_set" "tf-test-pxe-client" {
						name = "PXEClient"
						options = "next-server"
						depends_on = [
							"ros_dhcp_server_option.tf-test-next-server",
						]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerOptionSetExists("ros_dhcp_server_option_set.tf-test-pxe-client", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_option_set.tf-test-pxe-client", "name", "PXEClient"),
					resource.TestCheckResourceAttr("ros_dhcp_server_option_set.tf-test-pxe-client", "options", "next-server"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_dhcp_server_option" "tf-test-next-server" {
						code = 66
						name = "next-server"
						value = "'192.168.0.2'"
					}

					resource "ros_dhcp_server_option" "tf-test-bootfile" {
						code = 67
						name = "bootfile"
						value = "'pxelinux.0'"
					}

					resource "ros_dhcp_server_option_set" "tf-test-pxe-client" {
						name = "pxe-client"
						options = "next-server,bootfile"
						depends_on = [
							"ros_dhcp_server_option.tf-test-next-server",
							"ros_dhcp_server_option.tf-test-bootfile"
						]
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerOptionSetExists("ros_dhcp_server_option_set.tf-test-pxe-client", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_option_set.tf-test-pxe-client", "name", "pxe-client"),
					resource.TestCheckResourceAttr("ros_dhcp_server_option_set.tf-test-pxe-client", "options", "next-server,bootfile"),
				),
			},
		},
	})
}

func resourceDHCPServerOptionSetExists(n string, r *rc.ResourceDHCPServerOptionSet) resource.TestCheckFunc {
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

func testAccCheckDHCPServerOptionSetDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_dhcp_server_option" {
			continue
		}

		r := &rc.ResourceDHCPServerOptionSet{
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
