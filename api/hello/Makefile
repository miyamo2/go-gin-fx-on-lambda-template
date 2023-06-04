build-helloAPI:
	go mod download
	go mod tidy
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${outputdir}/hello github.com/miyamo2theppl/go-gin-fx-on-lambda-template/api/hello
	chmod 755 ${outputdir}/hello
	zip -j ${outputdir}/hello.zip ${outputdir}/hello