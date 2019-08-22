# riodemo

Demo rio 

## Prep

To get started we need a few docker images to use in our demo. Start by cloning this repo.  All command in this article assum you are in the project root.

* `make docker-build` - this will build two docker images that are used throught this demo. The images are:
  * `flower:yellowsun` - This image will listen on port 8080 and respond to a GET request by displaying ing 
    ```
      {
        "Variety": "sunflower",
        "Color": "yellow"
      }
    ```

  	docker build --build-arg COLOR=yellow --build-arg VARIETY=sunflower -t flower:yellowsun .
	docker build --build-arg COLOR=blue   --build-arg VARIETY=sunflower -t flower:bluesun .


```bash

docker build -t hi:blue .
docker build -t hi:green .
# test
docker run -it --rm -p 8080:80 hi:blue

# Install rio cli
$ curl -sfL https://get.rio.io | sh –

# Install rio into cluster
rio install

# Verify
rio info
```

## Auto scale a service

```bash
rio run --ports 8080/http --name hi-service hi:blue

# Verify cert in browser
rio ps

# Add load
hey -z 3m -c 40 <url of service endbpoint>

# watch scale up then down
watch rio ps

# Watch Kailai
rio -s ps
```

## Canary Deployment

```bash
# Stage a new version
rio stage --image=hi:green hi-service

# Notice a new URL was created for your staged service
# We can hit each revision directly if we want
# The service endpoint still returns "blue" becauess it is 100% weight
rio revision hi-service

# Promote green service.
# The traffic will be shifted gradually. By default we apply a 5% shift every 5 seconds
[terminal 1] watch rio ps
[terminal 2] hey -z 3m -c 40 <url of service endbpoint>
[terminal 3] rio promote hi-service:v<the green revision>
[browser] open kiali (set filter "name != hi-service")

# Manually adjusting weight between revisions
$ rio weight hi-service:v0=5% hi-service:v???=95%
```

## Serverless

```bash
# for show
[terminal 1] watch rio ps
[terminal 2] watch kubectl get pods
[terminal 3]
rio run -p 8080/http --name hi-service --scale=0-10 hi:blue
rio scale hi-service=1
rio scale hi-service=2
rio scale hi-service=3
curl <service endpoint url>
rio scale hi-service=0-3
curl <service endpoint url>
```


## Adding Router

## Monitoring



## backup

```bash
rio run -p 8080/http --name rdemo monachus/rancher-demo:yellow
rio scale default/rdemo=1-5
rio scale default/rdemo=3
rio stage --image=monachus/rancher-demo:pink default/rdemo:v1
rio weight --rollout-interval 2 default/rdemo:v1=25
rio promote --rollout-interval 2 default/rdemo:v1

$ siege -c10
https://www.linode.com/docs/tools-reference/tools/load-testing-with-siege/
```