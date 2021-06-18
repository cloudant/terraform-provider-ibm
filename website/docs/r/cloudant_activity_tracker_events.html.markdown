---
layout: "ibm"
page_title: "IBM : cloudant_activity_tracker_events"
description: |-
  Manages cloudant_activity_tracker_events.
subcategory: "Cloudant"
---

# ibm\_cloudant_activity_tracker_events

Provides a resource for cloudant_activity_tracker_events. This allows cloudant_activity_tracker_events to be created, updated and deleted.

## Example Usage

```hcl
resource "cloudant_activity_tracker_events" "cloudant_activity_tracker_events" {
}
```

## Argument Reference

The following arguments are supported:

* `types` - (Required, Forces new resource, List) An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. "management" is a required element of this array.
  * Constraints: The minimum length is `1` item.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_activity_tracker_events.

## Import

You can import the `cloudant_activity_tracker_events` resource by using `types`. An array of event types that are being sent to IBM Cloud Activity Tracker for the IBM Cloudant instance. Allowable values: management, data. "management" is a required element of this array.

```
$ terraform import cloudant_activity_tracker_events.cloudant_activity_tracker_events <types>
```
