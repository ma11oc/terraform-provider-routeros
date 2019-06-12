package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceDHCPServerOption(t *testing.T) {
	var r = &rc.ResourceDHCPServerOption{
		Code:  66,
		Name:  "next-server",
		Value: "'192.168.0.2'",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDHCPServerOptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprint(`
					resource "ros_dhcp_server_option" "tf-test-dhcp-server-option" {
						code = 66
						name = "next-server"
						value = "'192.168.0.2'"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerOptionExists("ros_dhcp_server_option.tf-test-dhcp-server-option", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_option.tf-test-dhcp-server-option", "code", "66"),
					resource.TestCheckResourceAttr("ros_dhcp_server_option.tf-test-dhcp-server-option", "name", "next-server"),
					resource.TestCheckResourceAttr("ros_dhcp_server_option.tf-test-dhcp-server-option", "value", "'192.168.0.2'"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_dhcp_server_option" "tf-test-dhcp-server-option" {
						code = 67
						name = "bootfile"
						value = "'pxelinux.0'"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDHCPServerOptionExists("ros_dhcp_server_option.tf-test-dhcp-server-option", r),
					resource.TestCheckResourceAttr("ros_dhcp_server_option.tf-test-dhcp-server-option", "code", "67"),
					resource.TestCheckResourceAttr("ros_dhcp_server_option.tf-test-dhcp-server-option", "name", "bootfile"),
					resource.TestCheckResourceAttr("ros_dhcp_server_option.tf-test-dhcp-server-option", "value", "'pxelinux.0'"),
				),
			},
		},
	})
}

func resourceDHCPServerOptionExists(n string, r *rc.ResourceDHCPServerOption) resource.TestCheckFunc {
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

func testAccCheckDHCPServerOptionDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_dhcp_server_option" {
			continue
		}

		r := &rc.ResourceDHCPServerOption{
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
