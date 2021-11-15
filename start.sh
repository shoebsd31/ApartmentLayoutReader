#!/bin/bash

docker build --target dev . -t go
docker run -it -v ${PWD}:/work go sh
