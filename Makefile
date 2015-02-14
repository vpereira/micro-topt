default: build

clean:
	rm -rf totp-cli

build:
	go build -o totp-cli
