init:
	glide up

build:
	go build -o cmd/main cmd/main.go

clean:
	rm -rf vendor/
	rm cmd/main
