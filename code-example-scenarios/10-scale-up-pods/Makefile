build:
	docker build --tag coffee-shop .

run:	build
	docker run -d --name coffee-shop coffee-shop

stop:	
	docker stop coffee-shop
	docker container rm coffee-shop

deploy:	
	kubectl apply -f deployment.yaml
	kubectl apply -f ingress.yaml

delete-deployment:
	kubectl delete -f deployment.yaml

up: run deploy

forward:
	kubectl port-forward deployment/coffee-shop 8080:8080

down: stop delete-deployment
		kubectl delete -f ingress.yaml