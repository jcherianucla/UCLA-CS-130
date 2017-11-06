#!/bin/bash
docker build -t grade-portal-api .
docker run -it --rm -p 3000:8080 grade-portal-api
