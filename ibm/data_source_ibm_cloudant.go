// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMCloudant() *schema.Resource {
	riSchema := dataSourceIBMResourceInstance().Schema

	riSchema["service"] = &schema.Schema{
		Description: "The service type of the instance",
		Type:        schema.TypeString,
		Computed:    true,
	}

	return &schema.Resource{
		Read:   dataSourceIBMCloudantRead,
		Schema: riSchema,
	}
}

func dataSourceIBMCloudantRead(d *schema.ResourceData, meta interface{}) error {
	err := dataSourceIBMResourceInstanceRead(d, meta)
	if err != nil {
		return err
	}

	err = setCloudantResourceControllerURL(d, meta)
	if err != nil {
		return err
	}

	return nil
}

func setCloudantResourceControllerURL(d *schema.ResourceData, meta interface{}) error {
	crn := d.Get(ResourceCRN).(string)
	rcontroller, err := getBaseController(meta)
	if err != nil {
		return err
	}
	d.Set(ResourceControllerURL, rcontroller+"/services/cloudantnosqldb/"+url.QueryEscape(crn))

	return nil
}
