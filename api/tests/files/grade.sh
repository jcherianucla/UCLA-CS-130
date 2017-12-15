#!/bin/bash
f=$1
name="project"
g++ -m32 -o $name $f

res=$(./$name)

if [ "$res" == 42 ]
then
	echo "Score: 100"
	echo "Success"
else
	echo "Score: 0"
	echo "Test1: Your program should output 42"
fi
