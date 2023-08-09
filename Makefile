build:
	go build -o main
	chmod +x ./main
	./main

bench:
	go test ./... -bench=. -count=8
