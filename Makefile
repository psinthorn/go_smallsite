# Start applicaton by run command below you can use start or run 
# use start to skip test 
# run migrate before run start
seed: 
	soda migrate
delete_seed:
	soda migrate down

start:
	go build -o go_smallsite app/*.go && ./go_smallsite

run:
	go run app/*.go


# Use docker
# Build docker image and start server use docker with compose
docker:
	docker-compose up

# Remark: for re-build docker image
#  - every time we edit or update Dockerfile we need to re-build docker image by running 
#  - docker-compose up --build 
docker-rebuild:
	docker-compose up --build