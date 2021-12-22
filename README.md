# dockerGoApp
docker simple build image using golang with CI/CD

build image using dockerfile :
```
docker build -t go-app .
```

running container :
```
docker run -dp 8080:8080 --name app go-app
```

test web app :
```
curl http://localhost:8080 
```
or open your browser
