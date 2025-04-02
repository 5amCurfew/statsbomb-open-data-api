data "google_iam_policy" "service_account_user" {
  binding {
    role = "roles/iam.serviceAccountUser"

    members = [
      "user:samueltobyknight@gmail.com",
    ]
  }
}

resource "google_service_account" "service_account_admin" {
  account_id    = "sa-admin"
  display_name  = "Admin SA"
  description   = "Service Account with Admin policy"
}

resource "google_service_account_iam_policy" "admin_account_iam" {
  service_account_id = google_service_account.service_account_admin.name
  policy_data        = data.google_iam_policy.admin.policy_data
}

resource "google_service_account" "service_account_reader" {
  account_id    = "sa-reader"
  display_name  = "Reader SA"
  description   = "Service Account with Reader policy"
}

resource "google_service_account_iam_policy" "reader_account_iam" {
  service_account_id = google_service_account.service_account_reader.name
  policy_data        = data.google_iam_policy.reader.policy_data
}