terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "4.42.1"
    }
  }
}

variable "project" {
  type    = string
  default = "extended-acumen-366113"
}

variable "zone" {
  type    = string
  default = "europe-west1-b"
}

provider "google" {
  project = var.project
  zone    = var.zone
}

resource "google_cloud_run_service" "backend" {
  name     = "backend"
  location = "europe-west1"

  template {
    spec {
      containers {
        image = "europe-west3-docker.pkg.dev/extended-acumen-366113/web-tech/backend"
        ports {
          container_port = 8080
        }
        volume_mounts {
          mount_path = "/secrets"
          name       = "Firebase"
        }
        resources {
          limits = {
            memory = "128Mi"
          }
        }
        env {
          name  = "PROJECT_ID"
          value = "web-tech-b4fa9"
        }
      }
      volumes {
        name = "Firebase"
        secret {
          secret_name = "Firebase"
          items {
            key  = "latest"
            path = "firebase.json"
          }
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "1"
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}


data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth_backend" {
  location = google_cloud_run_service.backend.location
  project  = google_cloud_run_service.backend.project
  service  = google_cloud_run_service.backend.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

resource "google_cloud_run_service" "frontend" {
  name     = "frontend"
  location = "europe-west1"
  template {
    spec {
      containers {
        image = "europe-west3-docker.pkg.dev/extended-acumen-366113/web-tech/frontend"
        ports {
          container_port = 80
        }
        resources {
          limits = {
            memory = "128Mi"
          }
        }
        env {
          name  = "VITE_API_URL"
          value = google_cloud_run_service.backend.status.0.url
        }
      }
    }
    metadata {
      annotations = {
        "autoscaling.knative.dev/maxScale" = "1"
      }
    }
  }
  traffic {
    percent         = 100
    latest_revision = true
  }
}

resource "google_cloud_run_service_iam_policy" "noauth_frontend" {
  location = google_cloud_run_service.frontend.location
  project  = google_cloud_run_service.frontend.project
  service  = google_cloud_run_service.frontend.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
