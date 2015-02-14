default: build

clean:
	rm -rf totp

build:
	go build -o totp
