.PHONY: run-server
run-server:
	gin run src/main.go

.PHONY: run-unocss
run-unocss:
	yarn dev

.PHONY: run
run:
	@make -j 2 run-unocss run-server
