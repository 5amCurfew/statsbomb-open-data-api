resource "google_storage_bucket" "data" {
  name          = "${var.project}-data"
  location      = var.default_region
  force_destroy = true

  lifecycle_rule {
    condition {
      age = 90
    }
    action {
      type = "Delete"
    }
  }
}