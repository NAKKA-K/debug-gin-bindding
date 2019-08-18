
run:
	go run main.go

curl:
	curl -X POST -H "Content-Type:application/json" -d '{"text":"hello"}' http://localhost:8080/post
