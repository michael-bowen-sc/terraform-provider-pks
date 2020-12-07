package tkgi

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"os"
	"testing"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"tkgi": testAccProvider,
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

type preCheckFunc = func(*testing.T)

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("TKGI_HOSTNAME"); v == "" {
		t.Fatal("TKGI_HOSTNAME must be set for acceptance tests")
	}
}

func generateRandomResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}
