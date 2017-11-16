#!/bin/bash
wget http://localhost:5000/pkg/github.com/jcherianucla/gradeportal/ --recursive --page-requisites --convert-links --no-parent --no-host-directories -e robots=off -P ./docs
