package routeros

import (
	"fmt"
	"testing"

	// FIXME
	rc "github.com/ma11oc/go-routerosclient"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceDNSStaticRecord(t *testing.T) {
	var r = &rc.ResourceDNSStaticRecord{
		Address:  "169.254.169.254",
		Disabled: false,
		Comment:  "managed_by_terraform",
		Name:     "some-host.example.tld",
		TTL:      "1d",
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDNSStaticRecordDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
					resource "ros_dns_static_record" "tf-test-host" {
						address = "169.254.169.254"
						name = "tf-test-host.example.tld"
						disabled = false
						ttl = "1d"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDNSStaticRecordExists("ros_dns_static_record.tf-test-host", r),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "address", "169.254.169.254"),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "name", "tf-test-host.example.tld"),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "disabled", "false"),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "ttl", "1d"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "ros_dns_static_record" "tf-test-host" {
						address = "169.254.169.253"
						name = "tf-test-host-mod.example.tld"
						disabled = true
						ttl = "1w"
					}
				`),
				Check: resource.ComposeTestCheckFunc(
					resourceDNSStaticRecordExists("ros_dns_static_record.tf-test-host", r),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "address", "169.254.169.253"),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "disabled", "true"),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "name", "tf-test-host-mod.example.tld"),
					resource.TestCheckResourceAttr("ros_dns_static_record.tf-test-host", "ttl", "1w"),
				),
			},
		},
	})
}

func resourceDNSStaticRecordExists(n string, r *rc.ResourceDNSStaticRecord) resource.TestCheckFunc {
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

func testAccCheckDNSStaticRecordDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*rc.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ros_dns_static_record" {
			continue
		}

		r := &rc.ResourceDNSStaticRecord{
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
