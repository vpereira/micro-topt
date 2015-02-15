default: build

clean:
	rm -rf totp-cli
	rm -rf services/login/login

deps:
	go get github.com/jinzhu/gorm
	go get github.com/mattn/go-sqlite3
	go get github.com/go-martini/martini
	go get github.com/martini-contrib/binding
	go get github.com/martini-contrib/render
	go get github.com/martini-contrib/sessions

build:
	cd services/login && go build -o login
	go build -o totp-cli
