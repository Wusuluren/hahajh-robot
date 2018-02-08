docker ps -a | awk '{system("docker kill "$1";docker rm "$1)}'

docker run --name some-mongo -p 27017:27017 -v /e/docker/hahajh-robot:/data/db -d mongo:3.4.11
