# Mux Rest API
This application provides the sample api using golang, gorilla-mux and mongo

### Requirements
This application requires:

* [go] ~> v1.15.6
* [mongodb] ~> v4.4

### Variable Environments
For this application to run correctly, it is necessary to create the environment variables below:

* MONGO_URI="**mongodb://localhost:27017**"
* MONGO_DBNAME="**your database name**"
* APP_PORT="**:your application port**"
### Dependecies
* github.com/gorilla/mux v1.8.0
* github.com/spf13/viper v1.7.1
* github.com/urfave/negroni v1.0.0
* go.mongodb.org/mongo-driver v1.5.0
