[![Testing](https://github.com/herlianto-github/Todos/actions/workflows/aws.yml/badge.svg?branch=testing)](https://github.com/herlianto-github/Todos/actions/workflows/aws.yml)

# Todos
Week 5 - Project Todo List

Todos REST API build using echo server.

The code implementation was inspired by [Layered Architecture](https://www.oreilly.com/library/view/software-architecture-patterns/9781491971437/ch01.html)

- **Configs**<br/>Contain database and http configuration
- **Delivery (API)**<br/>API http handlers or controllers
- **Entities**<br/>Contain database model
- **Repository**<br/>Contain implementation entities database anq query with ORM.
- **Utils**<br/>Contain database driver (mySQL)

# Setup Database
1. Read Documentation [link](https://www.digitalocean.com/community/tutorials/how-to-create-a-new-user-and-grant-permissions-in-mysql)
2. Create User Database in folder 'configs'

# Setup Apps Via Docker
1. Get into todos directory
2. Create docker network
    ```
    docker network create koneksiku
    ```
3. Docker network ip
    ```
    docker network inspect koneksiku
    ```
4. Create docker database image 'mySQL'
    ```
    docker run -d \
    --network koneksiku --network-alias mysql \
    -v todo-mysql-data:/var/lib/mysql \
    -e MYSQL_USER=todosadmin \
    -e MYSQL_ROOT_PASSWORD=todos123 \
    -e MYSQL_DATABASE=to_do_lists_test \
    mysql:latest
    ```
5. Get into docker mysql shell
    ```
    docker exec -it dockerid mysql -utodosadmin -ptodos123
    ```
6. Query something
7. Open New Terminal
8. Create docker image from 'Dockerfile'
    ```
    docker built -t imagename .
    ```
9. Run apps inside docker image
    ```
    docker run -it --network koneksiku imagename
    ```
    

# How To Run Apps
1. Create Database Based on Configs
2. Execute Command:
    ```console
    go run main.go
    ```
3. Read Documentation in folder `OpenApi`


# How To Test Apps
```
COMING SOON
```

