.PHONY: test
test:
	start=`date +%s` &&\
		go test `go list ./pkg/... | grep -v mocked/*` -p 1 &&\
	finish=`date +%s` &&\
	echo "Executed in ~$$((finish - start))s"

.PHONY: coverage
coverage:
	start=`date +%s` &&\
		go test `go list ./pkg/... | grep -v mocked/*` -p 1 -coverprofile=./coverage.out && go tool cover -html=coverage.out &&\
	finish=`date +%s` &&\
	echo "Executed in ~$$((finish - start))s"
