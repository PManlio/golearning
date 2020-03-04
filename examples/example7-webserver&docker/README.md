build docker image:

```
docker build -t img-name:optional-tag . -f Dockerfile-dev
```

run docker image:
```
docker run -p 3030:3001 -it img-name:optional-tag
```
