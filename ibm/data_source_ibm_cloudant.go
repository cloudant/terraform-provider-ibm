// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMCloudant() *schema.Resource {
	riSchema := dataSourceIBMResourceInstance().Schema

	riSchema["service"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "The service type of the instance",
	}

	riSchema["version"] = &schema.Schema{
		Type:        schema.TypeString,
		Computed:    true,
		Description: "Vendor version.",
	}

	riSchema["features"] = &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: "List of enabled optional features.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	riSchema["features_flags"] = &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: "List of feature flags.",
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	riSchema["include_data_events"] = &schema.Schema{
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Include data event types in events sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. By default only emitted events are of \"management\" type.",
	}

	riSchema["capacity"] = &schema.Schema{
		Type:        schema.TypeInt,
		Computed:    true,
		Description: "A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.",
	}

	riSchema["throughput"] = &schema.Schema{
		Type:        schema.TypeMap,
		Computed:    true,
		Description: "Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes.",
		Elem: &schema.Schema{
			Type: schema.TypeInt,
		},
	}

	riSchema["enable_cors"] = &schema.Schema{
		Type:        schema.TypeBool,
		Computed:    true,
		Description: "Boolean value to turn CORS on and off.",
	}

	riSchema["cors_config"] = &schema.Schema{
		Type:        schema.TypeList,
		Computed:    true,
		Description: "Configuration for CORS.",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_credentials": &schema.Schema{
					Type:        schema.TypeBool,
					Computed:    true,
					Description: "Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.",
				},
				"origins": &schema.Schema{
					Type:        schema.TypeList,
					Computed:    true,
					Description: "An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.",
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
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

	client, err := getCloudantClient(d, meta)
	if err != nil {
		return err
	}

	err = setCloudantServerInformation(client, d)
	if err != nil {
		return err
	}

	err = setCloudantActivityTrackerEvents(client, d)
	if err != nil {
		return err
	}

	err = setCloudantInstanceCapacity(client, d)
	if err != nil {
		return err
	}

	err = setCloudantInstanceCors(client, d)
	if err != nil {
		return err
	}

	return nil
}
