#!/bin/bash

docker-entrypoint.sh postgres &

/app/main
