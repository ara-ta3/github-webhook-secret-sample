GOVENDOR=$(shell which govendor)
HEROKU=$(shell which heroku)
token=
heroku_app_name=github-webhook-secret-sample
port=8080

run:
	env SECRET_TOKEN=$(token) \
		PORT=$(port) \
		go run main.go

install:
	$(GOVENDOR) install 

govendor:
	go get -u github.com/kardianos/govendor

deploy/heroku:
	git push heroku master

open:
	$(HEROKU) open --app $(heroku_app_name)

deploy/heroku/env:
	$(HEROKU) config:set --app $(heroku_app_name) SECRET_TOKEN=$(token)

deploy/heroku/setup: $(HEROKU)
	$(HEROKU) git:remote --app $(heroku_app_name)

$(HEROKU):
	which heroku || echo 'please install heroku cli https://devcenter.heroku.com/articles/heroku-cli'
