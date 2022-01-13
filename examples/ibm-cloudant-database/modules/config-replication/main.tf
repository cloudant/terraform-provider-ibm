
// Provision replication database
resource "ibm_cloudant_database" "cloudant_replicator_db" {
  instance_crn  = var.cloudant_instance_crn
  db            = "_replicator"
  partitioned   = var.cloudant_database_partitioned
  q             = var.cloudant_database_q
}

// Provision cloudant_replication resource instance
resource "ibm_cloudant_replication" "cloudant_replication_doc" {
  instance_crn  = ibm_cloudant_database.cloudant_replicator_db.instance_crn
  doc_id        = var.cloudant_replication_doc_id

  replication_document {
    id            = var.cloudant_replication_doc_id
    create_target = var.create_target
    continuous    = var.continuous
    cancel        = false

    source {
      auth {
        iam {
          api_key = var.source_api_key
        }
      }
      url = "https://${var.source_host}/${var.db_name}"
    }

    target {
      auth {
        iam {
          api_key = var.target_api_key
        }
      }
      url = "https://${var.target_host}/${var.db_name}"
    }
  }

  depends_on = [ibm_cloudant_database.cloudant_replicator_db]
}

data "ibm_cloudant_replication" "read_doc" {
  instance_crn  = ibm_cloudant_replication.cloudant_replication_doc.instance_crn
  doc_id        = var.cloudant_replication_doc_id
}