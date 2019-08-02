format:
	go fmt ./...

sonar:
	export PATH=$$PATH:/Users/yuzhipeng/Desktop/sonar-scanner-4.0.0.1744-macosx/bin
	sonar-scanner -Dsonar.projectKey=algorithm -Dsonar.sources=.

test:
	go test -cover -v ./golang/... -json -coverprofile=coverage-all.out > test.json

