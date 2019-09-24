docker-build:
	docker build --build-arg COLOR=yellow -t flower:yellow .
	docker build --build-arg COLOR=blue   -t flower:blue .
	docker build --build-arg COLOR=red   -t flower:red .

up:
	rio up --namespace n1 --file ./Riofile 
	#rio scale n1/add=1-5
	#rio scale mult=1-20
	#rio scale power=1-20

routs:
	rio route add add/to-add-v0 to n1/add

run:
	rio run --ports 8080/http --name hi-service --env FEATURE=green --version v1 hi:v1	
	

