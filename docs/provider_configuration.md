# TKGI Provider

For provisioning Kubernetes clusters through the TKGI API. The provider needs to be configured with the proper credentials before it can be used.

## Example Usage

```hcl
# Configure the TKGI provider.
provider "tkgi" {
  token  = "${var.token}"
  hostname = "${var.target}"
}

# Create a cluster
resource "tkgi_cluster" "example" {
  # ...
}
```

## Argument Reference

One of either `token`, or `client_id` + `client_secret` must be specified to authenticate with TKGI:

- `token` - (Optional) A Bearer token used to login to TKGI. This can be retrieved from the TKGI UAA with the following curl command: `BEARER_TOKEN="$(curl -s https://${TKGI_ADDRESS}:8443/oauth/token -k -XPOST -H 'Accept: application/json;charset=utf-8' -u "client_id:client_secret" -H 'Content-Type: application/x-www-form-urlencoded;charset=utf-8' -d 'grant_type=client_credentials' | jq -r .access_token)`, using client credentials such as the "UAA Management Admin Client" credential in the TKGI Tile. The token can also be passed to the provider with the `TKGI_TOKEN` shell environment variable.
- `client_id` - Can also be passed to the provider with the `TKGI_CLIENT_ID` shell environment variable.
- `client_secret` - Can also be passed to the provider with the `TKGI_CLIENT_SECRET` shell environment variable.

The following additional arguments are supported:

- `hostname` - (Required) Hostname of the TKGI API to connect to. Can also be passed to the provider with the `TKGI_HOSTNAME` shell environment variable.
- `skip_ssl_validation` - (Optional) Default `false`. Can also be passed to the provider with the `TKGI_SKIP_SSL_VALIDATION` shell environment variable.
- `max_wait_min` - (Optional) Length of time (in minutes) that the provider will wait for TKGI operations to complete. Default: 20. Can also be passed to the provider with the `TKGI_MAX_WAIT_MIN` shell environment variable.
- `wait_poll_interval_sec` - (Optional) Frequency of polling (in seconds) while waiting for TKGI operations to complete. Default: 10. Can also be passed to the provider with the `TKGI_WAIT_POLL_INTERVAL_SEC` shell environment variable.
