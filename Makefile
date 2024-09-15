# migrate-create:
# 	@read -p "Enter migration name: " name; \
# 	go run tools/create_migration.go -name $$name

# migrate-create:
# 	@bash -c 'read -p "Enter migration name: " name; \
# 	go run tools/create_migration.go -name $$name'

migrate-create:
	@set /p name="Enter migration name: " && go run tools/create_migration.go -name %name%

