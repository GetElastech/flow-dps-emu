
# Build All Targets
all: build

# Test API service
test:
	go test -v -tags=relic ./...

# Build production Docker containers
build: upstream
	docker build -t onflow.org/flow-dps-emu-dev --target build-env .
	docker build -t onflow.org/flow-dps-emu --target production .

# Run DPS emulator service in Docker
run: build
	docker run -t -i --name flow_dps_emu --rm -p 4900:9000 onflow.org/flow-dps-emu

# Run API service attached to Flow localnet network in Docker
backend: backend-start backend-stop

# Stop localnet flow tests
backend-stop:
	bash -c 'cd upstream/flow-go/integration/localnet; ! test -f ./docker-compose.nodes.yml || make stop'

# Run a Flow network in localnet in Docker
backend-start: upstream
	bash -c 'cd upstream/flow-go/integration/localnet && make init && make start'

# Tools that are needed for building
upstream:
	mkdir -p upstream
	# Install the latest flow
	git clone https://github.com/onflow/flow-go.git upstream/flow-go || true
	# bash -c 'cd upstream/flow-go && git checkout 03634c1406e86a40860d5561bb067bb7799b6073'
	bash -c 'cd upstream/flow-go && git reset --hard'
	# FIX: Temporary patch on current flow builds for DPS.
	# FIX: We should add a CI test for DPS compatibility there.
	bash -c 'cd upstream/flow-go && patch -p 1 <../../resources/buildfix'
	bash -c 'cd upstream/flow-go && make install-tools'

# Clean all unused images and containers
clean: backend-stop
	docker system prune -a -f
	rm -rf upstream/

