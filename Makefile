export KUBECONFIG=$(shell echo "$$(pwd)/.kubeconfig")
export KUBE_CONFIG_PATH=${KUBECONFIG}
export SKAFFOLD_DEFAULT_REPO=rain-cloud-local-registry.localhost:5000

.PHONY: start
start:
	k3d cluster start rain-cloud-local || true

.PHONY: stop
stop:
	k3d cluster stop rain-cloud-local || true

.PHONY: terraform
terraform: start
	terraform apply -auto-approve

.PHONY: dev
dev: terraform
	skaffold dev

.PHONY: k9s
k9s:
	k9s

.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2 run

.PHONY: generate
generate:
	go run pkg/cmd/generate/main.go
	sed -i "s/GroupsGetGroupUsersDetails200JSONResponse_Ocs_Data_Users_AdditionalProperties/interface{}/" pkg/api/provisioning_api/openapi.gen.go
	sed -i "s/UsersGetUsersDetails200JSONResponse_Ocs_Data_Users_AdditionalProperties/interface{}/" pkg/api/provisioning_api/openapi.gen.go

.PHONY: clean
clean:
	k3d cluster delete rain-cloud-local
	rm -f .kubeconfig *.tfstate*
