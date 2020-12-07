
provider "tkgi" {
  hostname = "${var.tkgi_api_dns_name}"
  token = "${var.token}"
  skip_ssl_validation = true
}

resource "tkgi_cluster" "example" {
  name = "example1"
  external_hostname = "${var.k8s_api_dns_name}"
  plan = "small"
  num_nodes = 1
//network_profile = "xyz" # unimplemented
}

variable "k8s_api_dns_name" {
  type = "string"
}

variable "tkgi_api_dns_name" {
  type = "string"
}

variable "token" {
  type = "string"
}
