FROM mcr.microsoft.com/dotnet/sdk:8.0 AS dotnet-build-env
WORKDIR /data

# Copy everything
COPY . .
# Build and publish a release

RUN dotnet publish /data/src/Server -c Release -o bin

# Build runtime image
FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /data
COPY --from=dotnet-build-env /data/bin/ .
COPY docker.env .env
#COPY --from=dotnet-build-env /data/bin/appsettings.json .
RUN true
#COPY docker.env .env
RUN true
ENTRYPOINT ["dotnet","NutritionTracker.Web.dll"]