package routeros

import (
	"os"
	"strings"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ros": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("ROS_DEFAULT_ADDR"); v == "" {
		t.Fatal("ROS_DEFAULT_ADDR must be set for acceptance tests")
	}

	if v := os.Getenv("ROS_DEFAULT_USERNAME"); v == "" {
		t.Fatal("ROS_DEFAULT_USERNAME must be set for acceptance tests")
	}

	if v := os.Getenv("ROS_DEFAULT_PASSWORD"); v == "" {
		t.Fatal("ROS_DEFAULT_PASSWORD must be set for acceptance tests")
	}
}

func testAccEnabled() bool {
	v := os.Getenv("TF_ACC")
	return v == "1" || strings.ToLower(v) == "true"
}
