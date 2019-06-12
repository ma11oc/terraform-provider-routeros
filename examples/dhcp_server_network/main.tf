provider "ros" {
  address = "127.0.0.1:8728"
  username = "vagrant"
  password = "vagrant"
}

resource "ros_dhcp_server_option" "tf-test-next-server" {
  code = 66
  name = "next-server"
  value = "'192.168.0.254'"
}

resource "ros_dhcp_server_network" "tf-test-dhcp-network" {
  address = "192.168.0.0/24"
  boot_file_name = "tftpboot/pxelinux.0"
  dhcp_option = "${ros_dhcp_server_option.tf-test-next-server.name}"
  domain = "example.com"
  dns_server = "192.168.0.1"
  gateway = "192.168.0.1"
  netmask = "24"
  next_server = "192.168.0.10"
  ntp_server = "192.168.0.10"
  wins_server = "192.168.0.10"
}
