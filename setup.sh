cd fixtures
docker-compose down
docker volume prune -f
docker-compose up -d
cd ..
go build
./fabric-go-sdk

