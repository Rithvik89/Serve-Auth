dev:
	go run cmd/*.go

create_db:
	docker run -d \
	--name primary_db \
	-p 5432:5432 \
	-e POSTGRES_USER=root \
	-e POSTGRES_PASSWORD=mypassword \
	postgres 

remove_db:
	docker rm -f primary_db

create_redis:
	docker run -d \
	--name primary_kvs \
	redis

remove_redis:
	docker rm -f primary_kvs