---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "unifi_network Resource - terraform-provider-unifi"
subcategory: ""
description: |-
  unifi_network manages WAN/LAN/VLAN networks.
---

# unifi_network (Resource)

`unifi_network` manages WAN/LAN/VLAN networks.

## Example Usage

```terraform
variable "vlan_id" {
  default = 10
}

resource "unifi_network" "vlan" {
  name    = "wifi-vlan"
  purpose = "corporate"

  subnet       = "10.0.0.1/24"
  vlan_id      = var.vlan_id
  dhcp_start   = "10.0.0.6"
  dhcp_stop    = "10.0.0.254"
  dhcp_enabled = true
}

resource "unifi_network" "wan" {
  name    = "wan"
  purpose = "wan"

  wan_networkgroup = "WAN"
  wan_type         = "pppoe"
  wan_ip           = "192.168.1.1"
  wan_egress_qos   = 1
  wan_username     = "username"
  x_wan_password   = "password"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the network.
- `purpose` (String) The purpose of the network. Must be one of `corporate`, `guest`, `wan`, or `vlan-only`.

### Optional

- `dhcp_dns` (List of String) Specifies the IPv4 addresses for the DNS server to be returned from the DHCP server. Leave blank to disable this feature.
- `dhcp_enabled` (Boolean) Specifies whether DHCP is enabled or not on this network.
- `dhcp_lease` (Number) Specifies the lease time for DHCP addresses in seconds. Defaults to `86400`.
- `dhcp_relay_enabled` (Boolean) Specifies whether DHCP relay is enabled or not on this network.
- `dhcp_start` (String) The IPv4 address where the DHCP range of addresses starts.
- `dhcp_stop` (String) The IPv4 address where the DHCP range of addresses stops.
- `dhcp_v6_dns` (List of String) Specifies the IPv6 addresses for the DNS server to be returned from the DHCP server. Used if `dhcp_v6_dns_auto` is set to `false`.
- `dhcp_v6_dns_auto` (Boolean) Specifies DNS source to propagate. If set `false` the entries in `dhcp_v6_dns` are used, the upstream entries otherwise Defaults to `true`.
- `dhcp_v6_enabled` (Boolean) Enable stateful DHCPv6 for static configuration.
- `dhcp_v6_lease` (Number) Specifies the lease time for DHCPv6 addresses in seconds. Defaults to `86400`.
- `dhcp_v6_start` (String) Start address of the DHCPv6 range. Used in static DHCPv6 configuration.
- `dhcp_v6_stop` (String) End address of the DHCPv6 range. Used in static DHCPv6 configuration.
- `dhcpd_boot_enabled` (Boolean) Toggles on the DHCP boot options. Should be set to true when you want to have dhcpd_boot_filename, and dhcpd_boot_server to take effect.
- `dhcpd_boot_filename` (String) Specifies the file to PXE boot from on the dhcpd_boot_server.
- `dhcpd_boot_server` (String) Specifies the IPv4 address of a TFTP server to network boot from.
- `domain_name` (String) The domain name of this network.
- `igmp_snooping` (Boolean) Specifies whether IGMP snooping is enabled or not.
- `internet_access_enabled` (Boolean) Specifies whether this network should be allowed to access the internet or not. Defaults to `true`.
- `ipv6_interface_type` (String) Specifies which type of IPv6 connection to use. Must be one of either `static`, `pd`, or `none`. Defaults to `none`.
- `ipv6_pd_interface` (String) Specifies which WAN interface to use for IPv6 PD. Must be one of either `wan` or `wan2`.
- `ipv6_pd_prefixid` (String) Specifies the IPv6 Prefix ID.
- `ipv6_pd_start` (String) Start address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
- `ipv6_pd_stop` (String) End address of the DHCPv6 range. Used if `ipv6_interface_type` is set to `pd`.
- `ipv6_ra_enable` (Boolean) Specifies whether to enable router advertisements or not.
- `ipv6_ra_preferred_lifetime` (Number) Lifetime in which the address can be used. Address becomes deprecated afterwards. Must be lower than or equal to `ipv6_ra_valid_lifetime` Defaults to `14400`.
- `ipv6_ra_priority` (String) IPv6 router advertisement priority. Must be one of either `high`, `medium`, or `low`
- `ipv6_ra_valid_lifetime` (Number) Total lifetime in which the address can be used. Must be equal to or greater than `ipv6_ra_preferred_lifetime`. Defaults to `86400`.
- `ipv6_static_subnet` (String) Specifies the static IPv6 subnet when `ipv6_interface_type` is 'static'.
- `multicast_dns` (Boolean) Specifies whether Multicast DNS (mDNS) is enabled or not on the network (Controller >=v7).
- `network_group` (String) The group of the network. Defaults to `LAN`.
- `network_isolation_enabled`: (Boolean) Specifies whether this network should be not allowed to access other local networks. Defaults to `false`.
- `site` (String) The name of the site to associate the network with.
- `subnet` (String) The subnet of the network. Must be a valid CIDR address.
- `vlan_id` (Number) The VLAN ID of the network.
- `wan_dhcp_v6_pd_size` (Number) Specifies the IPv6 prefix size to request from ISP. Must be between 48 and 64.
- `wan_dns` (List of String) DNS servers IPs of the WAN.
- `wan_egress_qos` (Number) Specifies the WAN egress quality of service. Defaults to `0`.
- `wan_gateway` (String) The IPv4 gateway of the WAN.
- `wan_gateway_v6` (String) The IPv6 gateway of the WAN.
- `wan_ip` (String) The IPv4 address of the WAN.
- `wan_ipv6` (String) The IPv6 address of the WAN.
- `wan_netmask` (String) The IPv4 netmask of the WAN.
- `wan_networkgroup` (String) Specifies the WAN network group. Must be one of either `WAN`, `WAN2` or `WAN_LTE_FAILOVER`.
- `wan_prefixlen` (Number) The IPv6 prefix length of the WAN. Must be between 1 and 128.
- `wan_type` (String) Specifies the IPV4 WAN connection type. Must be one of either `disabled`, `static`, `dhcp`, or `pppoe`.
- `wan_type_v6` (String) Specifies the IPV6 WAN connection type. Must be one of either `disabled`, `static`, or `dhcpv6`.
- `wan_username` (String) Specifies the IPV4 WAN username.
- `x_wan_password` (String) Specifies the IPV4 WAN password.

### Read-Only

- `id` (String) The ID of the network.

## Import

Import is supported using the following syntax:

```shell
# import from provider configured site
terraform import unifi_network.mynetwork 5dc28e5e9106d105bdc87217

# import from another site
terraform import unifi_network.mynetwork bfa2l6i7:5dc28e5e9106d105bdc87217

# import network by name
terraform import unifi_network.mynetwork name=LAN
```
