buildv1:
	# docker build -t add:v1 ./services/add
	# docker build -t mult:v1 ./services/mult
	# docker build -t power:v1 ./services/power
	docker build -t hi:tag1 ./services/hi

# pushv1:
# 	docker push miletwo/add:v1
# 	docker push miletwo/mult:v1
# 	docker push miletwo/power:v1

up:
	rio up --namespace n1 --file ./Riofile 
	#rio scale n1/add=1-5
	#rio scale mult=1-20
	#rio scale power=1-20

routs:
	rio route add add/to-add-v0 to n1/add

run:
	rio run --ports 80/http --name hi-service --env FEATURE=green --version v1 hi:v1	
	

