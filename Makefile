

mocks:
	@mockery -name="StorageManager" \
	-dir="$(shell pwd)/internal/web" \
	-output="$(shell pwd)/internal/web/mocks/"
	@echo 'Mocks has been generated...'