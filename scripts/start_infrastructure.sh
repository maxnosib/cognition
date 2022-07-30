cd docker
docker network create net_cognition
docker-compose up -d db_cognition
export PGPASSWORD=postgres
sleep 10
docker-compose exec db_cognition createdb -U postgres cognition
echo "инфраструктура запущена"