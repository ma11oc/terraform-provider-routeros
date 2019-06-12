package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceDHCPServer(t *testing.T) {
	var r = &rc.ResourceDHCPServer{
		Disabled:  false,
		Interface: "tf-test-br",
		Name:      "tf-test-dhcp-srv",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDHCPServerDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprint(`
					resource "ros_interface_bridge" "tf-test-br" {
						name = "tf-test-br"
						fast_forward = true
						disabled = false
					}

					resource "ros_dhcp_server" "tf-test-dhcp-srv" {
						name = "tf-test-dhcp-srv"
						interface = "${ros_interface_bridge.tf-test-br.name}"
						disabled = false
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerExists("ros_dhcp_server.tf-test-dhcp-srv", r),
					resource.TestCheckResourceAttr("ros_dhcp_server.tf-test-dhcp-srv", "disabled", "false"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_interface_bridge" "tf-test-br" {
						name = "tf-test-br"
						fast_forward = true
						disabled = false
					}

					resource "ros_dhcp_server" "tf-test-dhcp-srv" {
						name = "tf-test-dhcp-srv-mod"
						interface = "${ros_interface_bridge.tf-test-br.name}"
						disabled = true
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerExists("ros_dhcp_server.tf-test-dhcp-srv", r),
					resource.TestCheckResourceAttr("ros_dhcp_server.tf-test-dhcp-srv", "disabled", "true"),
					resource.TestCheckResourceAttr("ros_dhcp_server.tf-test-dhcp-srv", "name", "tf-test-dhcp-srv-mod"),
				),
			},
		},
	})
}

func resourceDHCPServerExists(n string, r *rc.ResourceDHCPServer) resource.TestCheckFunc {
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

func testAccCheckDHCPServerDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_dhcp_server" {
			continue
		}

		r := &rc.ResourceDHCPServer{
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
