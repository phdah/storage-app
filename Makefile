build:
	go build -o bin cmd/main.go

clean:
	@rm bin/* || echo "No binary file(s)"
