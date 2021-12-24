# Local environment

.PHONY: run
run:
	@docker-compose -f deployments/docker-compose.yml up -d --build
	@echo "Local environment is up."
	@echo

.PHONY: stop
stop:
	@docker-compose -f deployments/docker-compose.yml down
	@echo "Local environment is down."
	@echo

# Tools

.PHONY: grpcui
grpcui:
	@grpcui -proto api/hasher.proto -plaintext localhost:8090