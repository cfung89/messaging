#! /bin/bash

sudo docker pull postgres

sudo docker run --name msg-auth -e POSTGRES_PASSWORD=$1 -p 5432:5432 -d postgres

sudo docker exec msg-auth createdb -U postgres auth
