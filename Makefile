CONTROLLER_GEN=$(shell which controller-gen)
APISERVER_BOOT=$(shell which apiserver-boot)

generate: codegen
	find . -name "*.go" -exec gci write --Section Standard --Section Default --Section "Prefix(github.com/everoute/loadbalancer)" {} +

image-generate:
	docker buildx build -f build/images/generate/Dockerfile -t everoute/generate ./build/images/generate/ --load

docker-generate: image-generate
	$(eval WORKDIR := /go/src/github.com/everoute/loadbalancer)
	docker run --rm -iu 0:0 -w $(WORKDIR) -v $(CURDIR):$(WORKDIR) everoute/generate make generate

# Generate CRD manifests
manifests:
	$(CONTROLLER_GEN) crd paths="./pkg/apis/..." output:crd:dir=deploy/crds output:stdout

# Generate deepcopy, client, openapi codes
codegen: manifests
	$(APISERVER_BOOT) build generated --generator openapi --generator client --generator deepcopy  \
		--copyright hack/no-boilerplate.go.txt
# 		--api-versions loadbalancer/v1alpha1
