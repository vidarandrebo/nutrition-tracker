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

#dotnet ef migrations add $(name)Production --project src/Infrastructure --startup-project src/WebAPI --output-dir Migrations --context NpgsqlContext -v -- --environment Production
#dotnet ef migrations add $(name)Dev --project src/Infrastructure --startup-project src/Server --output-dir Migrations --context SqliteContext
migration:
	cp docker.env src/Server/.env
	dotnet ef migrations add $(name)Production --project src/Migrations.Postgres --startup-project src/Server -v -- --environment Production
	dotnet ef migrations add $(name)Dev --project src/Migrations.Sqlite --startup-project src/Server
	rm src/Server/.env
