terraform {
  required_providers {
    iphelpers = {
      source = "registry.terraform.io/ffddorf/iphelpers"
    }
  }
}

data "iphelpers_eui64_address" "link_local" {
  mac_address = "21:82:3c:60:5d:7c"
}

output "link_local" {
  value = data.iphelpers_eui64_address.link_local.ipv6_address
}

data "iphelpers_eui64_address" "global" {
  mac_address = "fb:48:72:63:6a:3c"
  prefix      = "2001:678:b7c:201::"
}

output "global" {
  value = data.iphelpers_eui64_address.global.ipv6_address
}

