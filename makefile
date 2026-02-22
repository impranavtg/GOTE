build:
	@go build -o gote .

run: build
	@./gote
