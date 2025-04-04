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

resource "google_storage_bucket_iam_binding" "storage_object_data_reader_binding" {
  bucket = google_storage_bucket.data.name
  role = "roles/storage.objectViewer"
  members = [
    "serviceAccount:${google_service_account.service_account_reader.email}"
  ]
}