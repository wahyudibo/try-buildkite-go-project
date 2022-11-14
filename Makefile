SERVICE_NAME := api
GENERATED_DOCKER_IMAGES := $(shell docker images | grep $(SERVICE_NAME) | awk '{print $$1}')

svc-up:
	docker-compose -p $(SERVICE_NAME) up --build -d

svc-test:
	docker-compose -p $(SERVICE_NAME) exec app ./scripts/test.sh

svc-down:
	docker-compose -p $(SERVICE_NAME) down -v --remove-orphans

svc-img-prune:
ifneq ($(strip $(GENERATED_DOCKER_IMAGES)),)
	docker rmi -f $(GENERATED_DOCKER_IMAGES)
	@echo "generated docker images succesfully deleted"
else
	@echo "generated docker images are not found"
endif