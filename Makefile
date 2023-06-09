export KUBECONFIG=$(shell echo "$$(pwd)/.kubeconfig")
export KUBE_CONFIG_PATH=${KUBECONFIG}
export SKAFFOLD_DEFAULT_REPO=rain-cloud-local-registry.localhost:5000

.PHONY: terraform
terraform:
	terraform apply -auto-approve

.PHONY: dev
dev: terraform
	skaffold dev

.PHONY: k9s
k9s:
	k9s

.PHONY: lint
lint: lint-go lint-protobuf

.PHONY: lint-go
lint-go:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2 run

.PHONY: lint-protobuf
lint-protobuf:
	(cd pkg/controller && go run github.com/bufbuild/buf/cmd/buf@v1.21.0 lint)

.PHONY: generate
generate: generate-apps generate-protobuf

.PHONY: generate-apps
generate-apps:
	go run pkg/cmd/generate-apps.go
	sed -i "s/GroupsGetGroupUsersDetails200JSONResponse_Ocs_Data_Users_AdditionalProperties/interface{}/" pkg/api/provisioning_api/openapi.gen.go
	sed -i "s/UsersGetUsersDetails200JSONResponse_Ocs_Data_Users_AdditionalProperties/interface{}/" pkg/api/provisioning_api/openapi.gen.go

.PHONY: generate-protobuf
generate-protobuf:
	(cd pkg/controller && go run github.com/bufbuild/buf/cmd/buf@v1.21.0 generate)

.PHONY: clean
clean:
	k3d cluster delete rain-cloud-local
	rm -f .kubeconfig *.tfstate*
