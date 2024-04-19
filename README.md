## Go Static Server

## TODO

- [] Include a way to customize 404 and 500 pages

## Publish

docker build . --file Dockerfile --tag ghcr.io/brownhounds/go-static:latest
docker push ghcr.io/brownhounds/go-static:latest
