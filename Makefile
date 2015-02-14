default: build

clean:
	rm -rf totp-cli
	rm -rf services/login/login

build:
	cd services/login && go build -o login
	go build -o totp-cli
