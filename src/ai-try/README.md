### MySQL Microservice Mutex

1. Download Docker & pull mysql latest image. `https://hub.docker.com/_/mysql`

2. Download test data for mysql. #https://dev.mysql.com/doc/employee/en/employees-installation.html

3. Extract the test data and keep it handy for editing.

4. Open a new terminal & Use the docker-compose file to start the mysql contaier.

5. Open new terminal and edit the employees.sql file in the test data folder.

6. Copy the test data to docker container. `docker cp test_db/ <CONTAINER_NAME>:/tmp/`

7. Change as per your case. The folder str & name depends on the previous step.
`
source /tmp/test_db/load_departments.dump;
`

8. Open container shell `docker exec -it <CONTAINER_NAME> /bin/bash`

9. Check that the /tmp folder has the copied files present.

10. Go into the /tmp folder and run the command to load data from test data to mysql. `mysql -uroot -p<Root_PSWD> -t < test_db/employees.sql`. Wait for the execution to complete.

11. Open mysql cli and verify data is populated. `mysql -uroot -p<Root_PSWD>`
