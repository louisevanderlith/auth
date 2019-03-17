# auth
Auth.APP will provide the visuals for Login and Registration.

## Run with Docker
* $ docker build -t avosa/auth:latest .
* $ docker rm AuthDEV
* $ docker run -d -e RUNMODE=DEV -p 8094:8094 --network mango_net --name AuthDEV avosa/auth:latest 
* $ docker logs AuthDEV
