
configure:
	sudo sysctl -w vm.max_map_count=262144
	sudo docker pull postgres
	sudo docker pull elasticsearch
	
run_containers:
	sudo docker run -p 9200:9200 -p 9300:9300 -v /home/lucas.miguel/elastic_data:/usr/share/elasticsearch/data -d elasticsearch
	sudo docker run -p 5432:5432 -v /home/lucas.miguel/postgres_data:/var/lib/postgresql/data -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=oi -d postgres
