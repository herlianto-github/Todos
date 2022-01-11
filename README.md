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

