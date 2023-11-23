
setup: postgres_docker

postgres_docker:
	docker run -d --name crud-postgres -e POSTGRES_PASSWORD=test -v ./postgres_data:/var/lib/postgresql/data -p 5432:5432 postgres
