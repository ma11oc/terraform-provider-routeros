## Preface
This provider doesn't cover all the features of routeros. It's just a proof-of-concept.

## Build
```bash
make
```

## Run
- run VM with routeros
    ```bash
    vagrant up --provider=virtualbox
    ```

- make basic config file
    ```bash
    mkdir -p /tmp/terraform/terraform.d/linux_amd64/plugins/

    cat > /tmp/terraform/main.tf << EOF
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
    EOF
    ```
- copy `terraform-provider-routeros` binary in terraform's plugin directory
    ```bash
    make install
    ```

- init and apply terraform config
    ```
    cd /tmp/terraform/
    terraform init
    terraform apply
    ```

## Testing
### Acceptance
It means you have to up vagrant box with routeros.
```
vagrant up --provider=virtualbox
make test
```

## License
MIT
