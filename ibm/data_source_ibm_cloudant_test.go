// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCloudantDataSource_basic(t *testing.T) {
	dataSourceName := "data.ibm_cloudant.instance"
	serviceName := fmt.Sprintf("terraform-test-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMCloudantDataSourceConfig(serviceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataSourceName, "name", serviceName),
					resource.TestCheckResourceAttr(dataSourceName, "service", "cloudantnosqldb"),
					resource.TestMatchResourceAttr(dataSourceName, ResourceControllerURL, regexp.MustCompile("services/cloudantnosqldb/crn%3A.+")),
					resource.TestCheckResourceAttr(dataSourceName, "audit_event_types.0", "management"),
					resource.TestCheckResourceAttr(dataSourceName, "capacity.0.throughput.0.query", "5"),
					resource.TestCheckResourceAttr(dataSourceName, "cors.0.enable_cors", "true"),
					resource.TestCheckResourceAttrSet(dataSourceName, "version"),
					resource.TestCheckResourceAttrSet(dataSourceName, "features.0"),
					resource.TestCheckResourceAttrSet(dataSourceName, "features_flags.0"),
				),
			},
		},
	})
}

func testAccCheckIBMCloudantDataSourceConfig(serviceName string) string {
	return fmt.Sprintf(`

	resource "ibm_resource_instance" "cloudant" {
	  name     = "%s"
	  service  = "cloudantnosqldb"
	  plan     = "lite"
	  location = "us-south"
	}

	data "ibm_cloudant" "instance" {
	  name     = ibm_resource_instance.cloudant.name
	}

	`, serviceName)
}
