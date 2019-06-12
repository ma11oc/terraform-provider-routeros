provider "ros" {
  address = "127.0.0.1:8728"
  username = "vagrant"
  password = "vagrant"
}

resource "ros_dns_static_record" "tf-test-host" {
  address = "169.254.169.254"
  name = "tf-test-host.example.tld"
  disabled = false
  ttl = "1d"
}
