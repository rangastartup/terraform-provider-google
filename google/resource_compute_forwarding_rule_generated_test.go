// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeForwardingRule_ForwardingRuleBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeForwardingRuleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeForwardingRule_ForwardingRuleBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_forwarding_rule.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeForwardingRule_ForwardingRuleBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_forwarding_rule" "default" {
  name       = "website-forwarding-rule-%s"
  target     = "${google_compute_target_pool.default.self_link}"
  port_range = "80"
}

resource "google_compute_target_pool" "default" {
  name = "website-target-pool-%s"
}
`, val, val,
	)
}
