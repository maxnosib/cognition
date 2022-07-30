# Запускать при первом развертывании
# export PGPASSWORD=postgres
# docker-compose exec db_cognition createdb -U postgres cognition
# make migration-init
make migration-up
./main

