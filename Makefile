rm:
	docker rm check_location
run:
	docker run --rm --name check_location -p 3005:3005 check_location:latest

build:
	docker build -t check_location:latest .