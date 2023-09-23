#! /bin/bash
psql -U andres -h localhost -c "DROP DATABASE bcatest"
psql -U andres -h localhost -c "CREATE DATABASE bcatest"
psql -U andres -h localhost -d bcatest -c "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

