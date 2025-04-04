variable project {default="statsbomb-open-data-api"}
variable default_region {default="europe-west2"}
variable default_zone {default="europe-west2-a"}

// terraform init --backend-config=backend.hcl

terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
    }
    google-beta = {
      source = "hashicorp/google-beta"
    }
  }
  required_version = ">= 0.14.8"
  backend "gcs" {}
}

provider "google" {
    project = var.project
    region = var.default_region
    zone = var.default_zone
}

data "google_iam_policy" "admin" {
  binding {
    role = "roles/admin"

    members = [
        "user:samueltobyknight@gmail.com",
        "serviceAccount:${google_service_account.service_account_admin.email}"
    ]
  }

  binding {
    role = "roles/iam.serviceAccountUser"
    members = ["user:samueltobyknight@gmail.com"]
  }
}