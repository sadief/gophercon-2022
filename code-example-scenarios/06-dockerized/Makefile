build:
	docker build --tag coffee-shop .

run:	build
	docker run -d --name coffee-shop -p 8080:8080 coffee-shop

stop:	
	docker stop coffee-shop
	docker container rm coffee-shop