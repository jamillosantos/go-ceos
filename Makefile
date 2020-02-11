

test:
	@go run github.com/onsi/ginkgo/ginkgo -r --randomizeAllSpecs --randomizeSuites --cover --trace --race --compilers=2 -coverprofile=coverage.txt -covermode=atomic

check-fmt:
	@test -z $(go fmt ./...)

check-vet:
	@go vet ./...

check-gocycle:
	@go run github.com/fzipp/gocyclo -over 25 ./

check-ineffassign:
	@go run github.com/jamillosantos/ineffassign -i ./.history ./

check-misspell:
	@go run github.com/client9/misspell/cmd/misspell ./

check: check-fmt check-vet check-vet check-ineffassign check-misspell
	@echo "\e[32mAll checks passed.\e[0m"