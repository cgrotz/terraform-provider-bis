package bis

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var thingsSolutionID string

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"bis": testAccProvider,
	}
}

func testAccPreCheck(t *testing.T) {
	if os.Getenv("CLIENT_ID") == "" && os.Getenv("CLIENT_SECRET") == "" {
		t.Fatal("CLIENT_ID and CLIENT_SECRET must be set for acceptance tests")
	}

	if os.Getenv("THINGS_SOLUTION_ID") == "" {
		t.Fatal("THINGS_SOLUTION_ID must be set for acceptance tests")
	} else {
		thingsSolutionID = os.Getenv("THINGS_SOLUTION_ID")
	}
}

func TestAccBISProvider_Namespace(t *testing.T) {
	namespace := "de.cgrotz.actions_name_space1"
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: nil,
		Steps: []resource.TestStep{
			{
				Config: testAccNamespaceResource(namespace),
				Check: resource.ComposeTestCheckFunc(
					testAccNamespaceResourceExists("bis_things_namespace.test_namespace"),
				),
			},
		},
	})
}

func testAccNamespaceResource(n string) string {
	return fmt.Sprintf(`
	resource "bis_things_namespace" "test_namespace" {
		solution_id = "%s"
		namespace = "%s"
	}
	`, os.Getenv("THINGS_SOLUTION_ID"), n)
}

func testAccNamespaceResourceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Can't find Namespace: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Namespace ID not set")
		}
		return nil
	}
}
