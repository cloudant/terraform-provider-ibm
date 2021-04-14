// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/cloudant-go-sdk/cloudantv1"
	"github.com/IBM/go-sdk-core/v5/core"
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
		Default:     true,
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

	riSchema["audit_event_types"] = &schema.Schema{
		Description: "An array of event types that are being sent to IBM Cloud Activity Tracker with LogDNA for the IBM Cloudant instance. \"management\" is a required element of this array.",
		Type:        schema.TypeList,
		Required:    true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	}

	riSchema["capacity"] = &schema.Schema{
		Description: "Detailed information about provisioned throughput capacity.",
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"throughput": &schema.Schema{
					Description: "Schema for detailed information about throughput capacity with breakdown by specific throughput requests classes.",
					Type:        schema.TypeList,
					Required:    true,
					MaxItems:    1,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"blocks": {
								Type:        schema.TypeInt,
								Required:    true,
								Description: "A number of blocks of throughput units. A block consists of 100 reads/sec, 50 writes/sec, and 5 global queries/sec of provisioned throughput capacity.",
							},
							"query": {
								Type:        schema.TypeInt,
								Computed:    true,
								Description: "Provisioned global queries capacity in operations per second.",
							},
							"read": {
								Type:        schema.TypeInt,
								Computed:    true,
								Description: "Provisioned reads capacity in operations per second.",
							},
							"write": {
								Type:        schema.TypeInt,
								Computed:    true,
								Description: "Provisioned writes capacity in operations per second.",
							},
						},
					},
				},
			},
		},
	}

	riSchema["cors"] = &schema.Schema{
		Description: "Detailed information about CORS configuration.",
		Type:        schema.TypeList,
		Required:    true,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"allow_credentials": {
					Description: "Boolean value to allow authentication credentials. If set to true, browser requests must be done by using withCredentials = true.",
					Type:        schema.TypeBool,
					Default:     true,
					Optional:    true,
				},
				"enable_cors": {
					Description: "Boolean value to turn CORS on and off.",
					Type:        schema.TypeBool,
					Default:     true,
					Optional:    true,
				},
				"origins": {
					Description: "An array of strings that contain allowed origin domains. You have to specify the full URL including the protocol. It is recommended that only the HTTPS protocol is used. Subdomains count as separate domains, so you have to specify all subdomains used.",
					Type:        schema.TypeList,
					Required:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
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

	client, err := getCloudantClient(d, meta)
	if err != nil {
		return err
	}

	if _, ok := d.GetOk("audit_event_types"); ok {
		err := updateCloudantInstanceAuditEventTypes(client, d)
		if err != nil {
			return fmt.Errorf("Error updating activity tracker events: %s", err)
		}
	}

	plan := d.Get("plan").(string)
	if plan != "lite" {
		if _, ok := d.GetOk("capacity"); ok {
			err := updateCloudantInstanceCapacity(client, d)
			if err != nil {
				return fmt.Errorf("Error retrieving capacity throughput information: %s", err)
			}
		}
	}

	if _, ok := d.GetOk("cors"); ok {
		err := updateCloudantInstanceCors(client, d)
		if err != nil {
			return fmt.Errorf("Error updating CORS settings: %s", err)
		}
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

	client, err := getCloudantClient(d, meta)
	if err != nil {
		return err
	}

	// if resource was imported, set missing legacy_credentials to default true
	// since we don't have this value exposed on the broker
	if _, ok := d.GetOkExists("legacy_credentials"); !ok {
		d.Set("legacy_credentials", true)
	}

	err = setCloudantInstanceAuditEventTypes(client, d)
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

func resourceIBMCloudantUpdate(d *schema.ResourceData, meta interface{}) error {
	d.Set("service", "cloudantnosqldb")
	err := resourceIBMResourceInstanceUpdate(d, meta)
	if err != nil {
		return err
	}

	client, err := getCloudantClient(d, meta)
	if err != nil {
		return err
	}

	if d.HasChange("audit_event_types") {
		err := updateCloudantInstanceAuditEventTypes(client, d)
		if err != nil {
			return fmt.Errorf("Error updating activity tracker events: %s", err)
		}
	}

	if d.HasChange("capacity") {
		err := updateCloudantInstanceCapacity(client, d)
		if err != nil {
			return fmt.Errorf("Error retrieving capacity throughput information: %s", err)
		}
	}

	if d.HasChange("cors") {
		err := updateCloudantInstanceCors(client, d)
		if err != nil {
			return fmt.Errorf("Error updating CORS settings: %s", err)
		}
	}

	return resourceIBMCloudantRead(d, meta)
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

func getCloudantClient(d *schema.ResourceData, meta interface{}) (*cloudantv1.CloudantV1, error) {

	extensions := d.Get("extensions").(map[string]interface{})
	_, ok := extensions["endpoints.public"]
	if !ok {
		return nil, fmt.Errorf("Missing endpoints.public in extensions")
	}
	endpoint := extensions["endpoints.public"].(string)

	session, err := meta.(ClientSession).BluemixSession()
	if err != nil {
		return nil, err
	}

	token := session.Config.IAMAccessToken
	token = strings.Replace(token, "Bearer ", "", -1)

	client, err := cloudantv1.NewCloudantV1(&cloudantv1.CloudantV1Options{
		Authenticator: &core.BearerTokenAuthenticator{
			BearerToken: token,
		},
		URL: "https://" + endpoint,
	})
	if err != nil {
		return nil, fmt.Errorf("Error occured while configuring Cloudant service: %q", err)
	}

	return client, nil
}

func setCloudantInstanceAuditEventTypes(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	activityTrackerEvents, err := readCloudantInstanceAuditEventTypes(client)
	if err != nil {
		return fmt.Errorf("Error retrieving activity tracker events: %s", err)
	}
	if activityTrackerEvents.Types != nil {
		d.Set("audit_event_types", activityTrackerEvents.Types)
	}
	return nil
}

func readCloudantInstanceAuditEventTypes(client *cloudantv1.CloudantV1) (*cloudantv1.ActivityTrackerEvents, error) {
	opts := client.NewGetActivityTrackerEventsOptions()

	activityTrackerEvents, response, err := client.GetActivityTrackerEvents(opts)
	if err != nil {
		log.Printf("[DEBUG] Error retrieving activity tracker events: %s\n%s", err, response)
	}
	return activityTrackerEvents, err
}

func updateCloudantInstanceAuditEventTypes(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	rawEventTypes := d.Get("audit_event_types").([]interface{})
	auditEventTypes := expandStringList(rawEventTypes)

	opts := client.NewPostActivityTrackerEventsOptions(auditEventTypes)

	_, response, err := client.PostActivityTrackerEvents(opts)
	if err != nil {
		log.Printf("[DEBUG] Error updating activity tracker events: %s\n%s", err, response)
	}
	return err
}

func setCloudantInstanceCapacity(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	capacityThroughputInformation, err := readCloudantInstanceCapacity(client)
	if err != nil {
		return fmt.Errorf("Error retrieving capacity throughput information: %s", err)
	}
	if capacityThroughputInformation.Current != nil && capacityThroughputInformation.Current.Throughput != nil {
		throughput := capacityThroughputInformation.Current.Throughput
		capacity := []map[string]interface{}{
			map[string]interface{}{
				"throughput": []map[string]interface{}{
					map[string]interface{}{
						"blocks": throughput.Blocks,
						"query":  throughput.Query,
						"read":   throughput.Read,
						"write":  throughput.Write,
					},
				},
			},
		}
		d.Set("capacity", capacity)
	}
	return nil
}

func readCloudantInstanceCapacity(client *cloudantv1.CloudantV1) (*cloudantv1.CapacityThroughputInformation, error) {
	opts := client.NewGetCapacityThroughputInformationOptions()

	capacityThroughputInformation, response, err := client.GetCapacityThroughputInformation(opts)
	if err != nil {
		log.Printf("[DEBUG] Error getting capacity throughput information: %s\n%s", err, response)
	}
	return capacityThroughputInformation, nil
}

func updateCloudantInstanceCapacity(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	capacity := d.Get("capacity").([]interface{})[0].(map[string]interface{})
	throughput := capacity["throughput"].([]interface{})[0].(map[string]interface{})
	blocks := int64(throughput["blocks"].(int))

	putOpts := client.NewPutCapacityThroughputConfigurationOptions(blocks)

	capacityThroughputInformation, response, err := client.PutCapacityThroughputConfiguration(putOpts)
	if err != nil {
		log.Printf("[DEBUG] Error updating capacity throughput: %s\n%s", err, response)
		return err
	}

	retryCount := 0
	getOpts := client.NewGetCapacityThroughputInformationOptions()
	current := *capacityThroughputInformation.Current.Throughput.Blocks
	target := *capacityThroughputInformation.Target.Throughput.Blocks

	for current != target {
		time.Sleep(200 * time.Millisecond)

		capacityThroughputInformation, response, err = client.GetCapacityThroughputInformation(getOpts)
		if err != nil {
			log.Printf("[DEBUG] Error retrieving capacity throughput information: %s\n%s", err, response)
			return err
		}

		current = *capacityThroughputInformation.Current.Throughput.Blocks
		target = *capacityThroughputInformation.Target.Throughput.Blocks
		retryCount++

		// wait up to 5 sec for capacity to sync up with target
		if retryCount > 25 {
			return fmt.Errorf("Retry count exceeded")
		}
	}

	return nil
}

func setCloudantInstanceCors(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	corsInformation, err := readCloudantInstanceCors(client)
	if err != nil {
		return fmt.Errorf("Error retrieving CORS config: %s", err)
	}
	if corsInformation != nil {
		cors := []map[string]interface{}{
			map[string]interface{}{
				"allow_credentials": corsInformation.AllowCredentials,
				"enable_cors":       corsInformation.EnableCors,
				"origins":           corsInformation.Origins,
			},
		}
		d.Set("cors", cors)
	}
	return nil
}

func readCloudantInstanceCors(client *cloudantv1.CloudantV1) (*cloudantv1.CorsInformation, error) {
	opts := client.NewGetCorsInformationOptions()

	corsInformation, response, err := client.GetCorsInformation(opts)
	if err != nil {
		log.Printf("[DEBUG] Error retrieving CORS config: %s\n%s", err, response)
	}
	return corsInformation, err
}

func updateCloudantInstanceCors(client *cloudantv1.CloudantV1, d *schema.ResourceData) error {
	cors := d.Get("cors").([]interface{})[0].(map[string]interface{})
	enableCors := cors["enable_cors"].(bool)
	allowCredentials := cors["allow_credentials"].(bool)
	origins := expandStringList(cors["origins"].([]interface{}))

	opts := client.NewPutCorsConfigurationOptions(origins)
	opts.SetEnableCors(enableCors)
	opts.SetAllowCredentials(allowCredentials)

	_, response, err := client.PutCorsConfiguration(opts)
	if err != nil {
		log.Printf("[DEBUG] Error updating CORS settings: %s\n%s", err, response)
	}
	return err
}
