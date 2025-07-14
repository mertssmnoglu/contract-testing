format:
    @gofmt -w ./

format-check:
    @gofmt -l ./

openapi-contract-test:
	@uvx schemathesis \
		--config-file ./schemathesis.toml \
		run ./openapi.yml \
		--url http://127.0.0.1:8080 \
		--report har \
		--report-dir .schemathesis-report

clean-reports:
	@rm -rf .hypothesis .schemathesis-report
