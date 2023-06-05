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
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.2 run

.PHONY: generate
generate:
	for app in core; do \
		go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0 -generate types,chi-server,strict-server -package $${app}api -o pkg/api/$${app}/openapi.gen.go pkg/api/$${app}/openapi.json; \
	done

.PHONY: clean
clean:
	k3d cluster delete rain-cloud-local
	rm -f .kubeconfig *.tfstate*
