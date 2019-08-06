CONTAINER_REPO_URL := release-package

buildv1:
	docker build -t miletwo/add-svr:v1 ./services/add
	docker build -t miletwo/mult-svr:v1 ./services/mult
	docker build -t miletwo/power-svr:v1 ./services/power

pushv1:
	docker push miletwo/add-svr:v1
	docker push miletwo/mult-svr:v1
	docker push miletwo/power-svr:v1

up:
	rio up --namespace m2 --file ./Riofile 
	#rio scale add-svr=1-20
	#rio scale mult-svr=1-20
	#rio scale power-svr=1-20

run:
	# rio -namespace m2 run -p 80/http --name add-svr miletwo/add-svr:v1.1	
