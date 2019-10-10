format:
	go fmt ./...

test: format
	go vet ./...
	go test -cover -v ./golang/... -json -coverprofile=coverage-all.out > test.json

sonar: test
	export PATH=$$PATH:/Users/yuzhipeng/Desktop/sonar-scanner-4.0.0.1744-macosx/bin
	sonar-scanner -Dsonar.projectKey=algorithm -Dsonar.sources=.

upload:
	git add .
	git commit -m "update load code"
	git push origin master


