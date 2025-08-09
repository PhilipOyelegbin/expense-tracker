swag:
	export PATH=$PATH:~/go/bin && swag init

run:
	go run main.go

build:
	go build -o ./app

tidy:
	go mod tidy