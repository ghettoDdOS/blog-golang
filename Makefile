.PHONY: run-db-server
run-db-server:
	docker compose up -d

.PHONY: run-server
run-server:
	air

.PHONY: run-unocss
run-unocss:
	yarn dev

.PHONY: run
run:
	@make run-db-server
	@make -j 2 run-unocss run-server
