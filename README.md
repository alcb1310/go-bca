# BUDGET CONTROL APPLICATION

This is the project in which the BCA backend app will be developed in.

## Techincal stack

![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

## Main Features

- Create projects to manage
- Add budget items to the project
- Add invoices, each invoice will decrease the values from the budget

## Environment Variables

In order to configure this project, the following environment variables are required:

```bash
DB_USERNAME
DB_HOST
DB_PASSWORD
DB_NAME
DB_PORT

PORT
```

## API Routes

- **/login** will allow the registered user to login to the application, on success it will return a JWT Token as a prove of authentication
