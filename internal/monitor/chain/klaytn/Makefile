GO  = GO111MODULE=on go

help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## make solo: build plugin (make plugin type= <solo>)
plugin:
	@mkdir -p ../../../../build
	$(GO) build --buildmode=plugin -o ../../../../build/klaytn.so *.go
