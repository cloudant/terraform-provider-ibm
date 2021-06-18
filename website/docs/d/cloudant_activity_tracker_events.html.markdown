---
layout: "ibm"
page_title: "IBM : cloudant_activity_tracker_events"
description: |-
  Get information about cloudant_activity_tracker_events
subcategory: "Cloudant"
---

# ibm\_cloudant_activity_tracker_events

Provides a read-only data source for cloudant_activity_tracker_events. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "cloudant_activity_tracker_events" "cloudant_activity_tracker_events" {
}
```

## Argument Reference

The following arguments are supported:


## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_activity_tracker_events.
* `types` - An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. "management" is a required element of this array.

