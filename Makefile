.PHONY: run
run:
	@docker-compose -f deployments/docker-compose.yml up -d --build

.PHONY: grpcui
grpcui:
	@grpcui -proto api/hasher.proto -plaintext localhost:8090