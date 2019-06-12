provider "ros" {
  address = "127.0.0.1:8728"
  username = "vagrant"
  password = "vagrant"
}

resource "ros_interface_bridge" "bridge0" {
    name = "bridge0"
    fast_forward = true
    disabled = false
}

resource "ros_dhcp_server" "dhcp0" {
    name = "dhcp0"
    interface = "${ros_interface_bridge.bridge0.name}"
}

resource "ros_dhcp_server_lease" "sample-host" {
    address = "169.254.169.254"
    server = "${ros_dhcp_server.dhcp0.name}"
    mac = "00:11:22:33:44:55"
    disabled = false
    client_id = "sample-host"
}

