provider "ros" {
  address = "127.0.0.1:8728"
  username = "vagrant"
  password = "vagrant"
}

resource "ros_dhcp_server_option" "tf-test-next-server" {
  code = 66
  name = "next-server"
  value = "'192.168.0.2'"
}

resource "ros_dhcp_server_option" "tf-test-bootfile" {
  code = 67
  name = "bootfile"
  value = "'pxelinux.0'"
}

resource "ros_dhcp_server_option_set" "tf-test-dhcp-server-option-set" {
  name = "pxe-client"
  options = "next-server,bootfile"
  depends_on = [
    "ros_dhcp_server_option.tf-test-next-server",
    "ros_dhcp_server_option.tf-test-bootfile"
  ]
}
