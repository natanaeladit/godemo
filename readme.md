##How to build docker file
```
docker build -t natanaeladit/godemo .
docker run -p 6060:8080 --rm natanaeladit/godemo
```

##Explore docker image
```
docker run --rm -it --entrypoint=/bin/bash natanaeladit/godemo
```

##Explore a running docker container
```
docker exec -it natanaeladit/godemo bash
```