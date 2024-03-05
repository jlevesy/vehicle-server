.PHONY: unit_test
unit_test:
	go test -count=1 -v ./...

.PHONY: integration_test
integration_test:
	go test -v -count=1 --tags=integration ./app
