clean:
	rm -rf src/Client/node_modules
	rm -rf src/Client/dist
	rm -rf src/Server/wwwroot/*
	rm -rf src/Server/Data
	find ./src/Infrastructure -type d \( -name "bin" -o -name "obj" \) -exec rm -rf {} +
	find ./src/Domain -type d \( -name "bin" -o -name "obj" \) -exec rm -rf {} +
	find ./src/Application -type d \( -name "bin" -o -name "obj" \) -exec rm -rf {} +
	find ./src/Server -type d \( -name "bin" -o -name "obj" \) -exec rm -rf {} +
	find ./src/Migrations.Sqlite -type d \( -name "bin" -o -name "obj" \) -exec rm -rf {} +
	find ./src/Migrations.Postgres -type d \( -name "bin" -o -name "obj" \) -exec rm -rf {} +

migration:
	cp docker.env src/Server/.env
	dotnet ef migrations add $(name)Postgres --project src/Migrations.Postgres --startup-project src/Server -v -- --environment Production
	dotnet ef migrations add $(name)Sqlite --project src/Migrations.Sqlite --startup-project src/Server
	rm src/Server/.env
