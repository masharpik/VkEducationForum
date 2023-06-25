.PHONY: build generate run rerun stop dbrun dbstop dbdown dbclean dbrerun servrun functests buildtests perftests

port=5001

generate:
	swagger generate spec -o ./swagger.yml -m
	go generate ./...

build:
	docker build -t forum .

run:
	docker run --rm --memory 2G --log-opt max-size=5M --log-opt max-file=3 --name forum -p ${port}:${port} -p 5432:5432 -t forum

rerun: build run

stop:
	docker stop forum

dbrun:
	docker-compose up

dbstop:
	docker-compose stop

dbdown:
	docker-compose down

dbclean:
	docker volume rm forumvkeducation_postgres_data

dbrerun: dbdown dbclean dbrun

servrun:
	go run forum/main.go

functests:
	./technopark-dbms-forum func -u http://localhost:${port}/api -r report.html

buildtests:
	go get -u -v github.com/mailcourses/technopark-dbms-forum@master
	go build github.com/mailcourses/technopark-dbms-forum
	go mod tidy

perftests:
	./technopark-dbms-forum fill --url=http://localhost:${port}/api --timeout=900
	./technopark-dbms-forum perf --url=http://localhost:${port}/api --duration=600 --step=60
