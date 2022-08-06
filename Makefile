
# Build All Targets
all: build

# Test API service
test:
	go test -v -tags=relic ./...

# Build production Docker containers
build:
	docker build -t onflow.org/flow-dps-emu --target build-env .

# Run DPS emulator service in Docker
run: build
	docker run -t -i --name flow_dps_emu --rm -p 4900:9000 onflow.org/flow-dps-emu .

# Run API service attached to Flow localnet network in Docker
backend: backend-start backend-stop

# Stop localnet flow tests
backend-stop:
	bash -c 'cd upstream/flow-go/integration/localnet && make stop'

# Run a Flow network in localnet in Docker
backend-start:
	# We might want to use testnet
	git clone https://github.com/onflow/flow-go.git upstream/flow-go || true
	# bash -c 'cd upstream/flow-go && git checkout 03634c1406e86a40860d5561bb067bb7799b6073'
	bash -c 'cd upstream/flow-go && make install-tools'
	bash -c 'cd upstream/flow-go/integration/localnet && make init && make start'

# Clean all unused images and containers
clean:
	docker system prune -a -f

