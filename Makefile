.PHONY: krew-wasm-demo
krew-wasm-demo: clear
	@go run . -0

.PHONY: clear
clear:
	clear
