.PHONY: run-server
run-server:
	gin run main.go

.PHONY: run-unocss
run-unocss:
	yarn dev

.PHONY: run
run:
	@make -j 2 run-unocss run-server
