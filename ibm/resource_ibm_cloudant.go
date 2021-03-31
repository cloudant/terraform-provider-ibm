// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
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

	riSchema["legacy_credentials"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		ForceNew:    true,
		Description: "Use both legacy credentials and IAM for authentication",
	}

	riSchema["environment_crn"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
		Description: "CRN of the IBM Cloudant Dedicated Hardware plan instance",
	}

	riSchema["cluster_location"] = &schema.Schema{
		Type:        schema.TypeString,
		Optional:    true,
		ForceNew:    true,
		Description: "The actual physical location of the Dedicated Hardware plan instance",
	}

	riSchema["hipaa"] = &schema.Schema{
		Type:        schema.TypeBool,
		Optional:    true,
		Description: "Instance is HIPAA ready in US locations",
	}

	riSchema["kms_instance_crn"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		RequiredWith: []string{"kms_key_crn"},
		Description:  "CRN of the Key Protect instance housing the encryption key for BYOK",
	}

	riSchema["kms_key_crn"] = &schema.Schema{
		Type:         schema.TypeString,
		Optional:     true,
		ForceNew:     true,
		RequiredWith: []string{"kms_instance_crn"},
		Description:  "CRN of the encryption key that is stored in the Key Protect instance",
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

	params := make(map[string]interface{})

	if legacyCredentials, ok := d.GetOkExists("legacy_credentials"); ok {
		params["legacyCredentials"] = fmt.Sprintf("%t", legacyCredentials)
	}

	if environmentCRN, ok := d.GetOk("environment_crn"); ok {
		params["environment_crn"] = environmentCRN
	}

	if clusterLocation, ok := d.GetOk("cluster_location"); ok {
		params["location"] = clusterLocation
	}

	if hipaa, ok := d.GetOkExists("hipaa"); ok {
		params["hipaa"] = fmt.Sprintf("%t", hipaa)
	}

	if kmsInstanceCRN, ok := d.GetOk("kms_instance_crn"); ok {
		params["kms_instance_crn"] = kmsInstanceCRN
	}

	if kmsKeyCRN, ok := d.GetOk("kms_key_crn"); ok {
		params["kms_key_crn"] = kmsKeyCRN
	}

	// copy values from "parameters" to params, unless they are already defined
	parameters, ok := d.GetOk("parameters")
	if ok {
		temp := parameters.(map[string]interface{})
		for k, v := range temp {
			if override, ok := params[k]; ok && override != v {
				log.Printf("[WARN] Overriding %q in 'parameters' to %s", k, override)
				continue
			}
			params[k] = v
		}
	}

	if len(params) > 0 {
		d.Set("parameters", params)
	}

	err := resourceIBMResourceInstanceCreate(d, meta)
	if err != nil {
		return err
	}

	// return original parameters on state
	d.Set("parameters", parameters)

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
