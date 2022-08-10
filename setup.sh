#! /bin/sh

docker build . -t llm
docker run -it llm sh