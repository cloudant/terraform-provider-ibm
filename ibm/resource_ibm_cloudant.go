// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMCloudant() *schema.Resource {
	riSchema := resourceIBMResourceInstance().Schema

	riSchema["service"] = &schema.Schema{
		Description: "The service type of the instance",
		Type:        schema.TypeString,
		Computed:    true,
	}

	return &schema.Resource{
		Create:   resourceIBMCloudantCreate,
		Read:     resourceIBMCloudantRead,
		Update:   resourceIBMCloudantUpdate,
		Delete:   resourceIBMResourceInstanceDelete,
		Exists:   resourceIBMResourceInstanceExists,
		Importer: &schema.ResourceImporter{},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
				return resourceTagsCustomizeDiff(diff)
			},
		),

		Schema: riSchema,
	}
}

func resourceIBMCloudantCreate(d *schema.ResourceData, meta interface{}) error {
	d.Set("service", "cloudantnosqldb")
	err := resourceIBMResourceInstanceCreate(d, meta)
	if err != nil {
		return err
	}

	return resourceIBMCloudantRead(d, meta)
}

func resourceIBMCloudantRead(d *schema.ResourceData, meta interface{}) error {
	err := resourceIBMResourceInstanceRead(d, meta)
	if err != nil {
		return err
	}

	err = setCloudantResourceControllerURL(d, meta)
	if err != nil {
		return err
	}

	return nil
}

func resourceIBMCloudantUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Set("service", "cloudantnosqldb")
	err := resourceIBMResourceInstanceUpdate(d, meta)
	if err != nil {
		return err
	}

	return resourceIBMCloudantRead(d, meta)
}
