package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceInterfaceBridge(t *testing.T) {
	var r = &rc.ResourceInterfaceBridge{
		Name:        "tf-test-br",
		FastForward: true,
		Disabled:    false,
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceBridgeDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "ros_interface_bridge" "tf-test-br" {
						name = "tf-test-br"
						fast_forward = true
						disabled = false
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceInterfaceBridgeExists("ros_interface_bridge.tf-test-br", r),
					resource.TestCheckResourceAttr("ros_interface_bridge.tf-test-br", "disabled", "false"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_interface_bridge" "tf-test-br" {
						name = "tf-test-br-mod"
						fast_forward = true
						disabled = true
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceInterfaceBridgeExists("ros_interface_bridge.tf-test-br", r),
					resource.TestCheckResourceAttr("ros_interface_bridge.tf-test-br", "disabled", "true"),
					resource.TestCheckResourceAttr("ros_interface_bridge.tf-test-br", "name", "tf-test-br-mod"),
				),
			},
		},
	})
}

func resourceInterfaceBridgeExists(n string, r *rc.ResourceInterfaceBridge) resource.TestCheckFunc {
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

func testAccCheckInterfaceBridgeDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_interface_bridge" {
			continue
		}

		r := &rc.ResourceInterfaceBridge{
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
