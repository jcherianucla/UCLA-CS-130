#!/bin/bash
docker build -t grade-portal .
docker run -it -p 5000:5000 grade-portal
