TF_PLUGINS_DIR := /tmp/terraform/terraform.d/plugins/linux_amd64/

default: build

build:
	go build

# install:
#   go install

install:
	install -m 0755 terraform-provider-routeros $(TF_PLUGINS_DIR)/terraform-provider-ros

test:
	TF_ACC=true \
	ROS_DEFAULT_ADDR=127.0.0.1:8728 \
	ROS_DEFAULT_USERNAME=vagrant \
	ROS_DEFAULT_PASSWORD=vagrant \
		go test -v -covermode=count -coverprofile=profile.cov ./routeros

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: build install test vet
