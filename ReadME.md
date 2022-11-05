#Golang kafka example

____

## How to run project

- docker-compose build --no-cache
- docker-compose up -d
- docker-compose ps
- docker-compose logs -f {service-name}
- send POST request to API' , http://127.0.0.1:8080/api/v1/kafka/push , for pushing messages 
- docker-compose down (remove and stop containers)

```
Body
{
    "id":2,
    "fio":"Sanzhar Anarbay",
    "group":"Group-1",
    "major":"Information Systems",
    "gpa":3.45
}

```
____

## Project Details
- visit http://localhost:9000/ (Portainer) , docker containers GUI
- open producer-api container logs, see the actions when send request to push Message into topic
- open consumer container logs, see the logs of the consuming messages from the topic

