provider "google" {
  project = "${var.projectid}"
  region = "${var.region}"
  credentials = "${file("${path.cwd}/${var.service_account_key_path}")}"
}

variable "projectid" { }
variable "region" { }
variable "service_account_key_path" { }
variable "zones" {
  type = "list"
  default = ["us-central1-a", "us-central1-b", "us-central1-c"]
}
variable "prefix" {
    type = "string"
    default = "cfcr"
}

resource "google_compute_global_address" "static_address" {
  name = "${var.prefix}-https-lb-address"
}

resource "google_dns_record_set" "a" {
  name = "${var.prefix}-gcp-lb.kubo.sh."
  managed_zone = "kubosh"
  type = "A"
  ttl  = 300

  rrdatas = ["${google_compute_global_address.static_address.address}"]
}

resource "google_compute_instance_group" "default" {
  // Count based on number of AZs
  count       = "3"
  name        = "${var.prefix}-httpslb-${element(var.zones, count.index)}"
  description = "terraform generated instance group that is multi-zone for https loadbalancing"
  zone        = "${element(var.zones, count.index)}"

  named_port {
    name = "kubernetes-master"
    port = "8443"
  }
}

resource "google_compute_target_https_proxy" "default" {
  name             = "${var.prefix}-cfcr-proxy"
  url_map          = "${google_compute_url_map.default.self_link}"
  ssl_certificates = ["https://www.googleapis.com/compute/v1/projects/cf-pcf-kubo/global/sslCertificates/kubo-sh-cert"]
}

resource "google_compute_url_map" "default" {
  name        = "${var.prefix}-https-lb"
  description = "a description"

  default_service = "${google_compute_backend_service.default.self_link}"
}

resource "google_compute_backend_service" "default" {
  name        = "${var.prefix}-k8s-master-service"
  port_name   = "kubernetes-master"
  protocol    = "HTTPS"
  timeout_sec = 60
  enable_cdn  = false

  backend {
    group = "${google_compute_instance_group.default.0.self_link}"
  }

  backend {
    group = "${google_compute_instance_group.default.1.self_link}"
  }

  backend {
    group = "${google_compute_instance_group.default.2.self_link}"
  }

  health_checks = ["${google_compute_https_health_check.default.self_link}"]
}

resource "google_compute_https_health_check" "default" {
  name               = "${var.prefix}-k8s-master"
  request_path       = "/healthz"
  port               = 8443
  check_interval_sec = 2
  timeout_sec        = 2
}

resource "google_compute_global_forwarding_rule" "cf-https" {
  name       = "${var.prefix}-cfcr-lb-https"
  ip_address = "${google_compute_global_address.static_address.address}"
  target     = "${google_compute_target_https_proxy.default.self_link}"
  port_range = "443"
}

output "ip_address" {
  value = "${google_compute_global_address.static_address.address}"
}
