# MERCHANT GO API

---

This project implements a simple Backend API for managing merchant and bank interactions, including customer login, payment, and logout functionalities. The API also logs activities to a history file. It is written in Go (Golang) and uses JSON files for simulating customer, merchant, and history data.

### FEATURES

1.**Login**

  A customer can log in if they exist in the system. If the customer does not exist, they will be rejected.

2.**Logout**

  A logged-in customer can log out, terminating their session.

3.**Payment**

  A logged-in customer can make a payment. There are no minimum or maximum transfer limits. Transfers are only allowed between registered customers.

4.**Activity Loggging**

  All activities (login, payment, logout) are logged to a history JSON file.

5. **Transaction Rollback Feature**

   If an error or issue occurs during a payment process, the transaction can be rolled back. This feature ensures that if a transaction fails, changes made to customer data and transaction history are reverted to their original state, maintaining system consistency and preventing loss of data or unauthorized balances.

---
## Project Structure
```
.
├── cmd
│   └── app
│       └── main.go
├── data
│   ├── customer.json
│   ├── customer.json.backup
│   ├── merchant.json
│   ├── merchant.json.backup
│   ├── merchant_sample.json
│   ├── transaction.json
│   ├── transaction.json.backup
│   └── transaction_sample.json
├── gin.log
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── dto
│   │   │   ├── CommonResponse.go
│   │   │   ├── Request.go
│   │   │   └── Response.go
│   │   ├── handler
│   │   │   ├── Auth.go
│   │   │   └── TransactionHandler.go
│   │   ├── models
│   │   │   ├── Customer.go
│   │   │   ├── Merchant.go
│   │   │   └── Transaction.go
│   │   ├── pkg
│   │   │   ├── data
│   │   │   │   └── Rollback.go
│   │   │   └── token
│   │   │       └── Token.go
│   │   ├── repositories
│   │   │   ├── CustomerRepository.go
│   │   │   ├── MerchantRepository.go
│   │   │   └── TransactionRepository.go
│   │   ├── router
│   │   │   └── Router.go
│   │   └── services
│   │       ├── AuthService.go
│   │       └── TransactionService.go
│   └── utils
│       ├── AuthUtil.go
│       └── JsonUtils.go
├── README.md
└── tests
    ├── e2e
    ├── integration
    ├── testdata
    │   ├── customer_sample.json
    │   ├── merchant_sample.json
    │   ├── merchant_sample.json.backup
    │   └── transaction_sample.json
    └── unit
        ├── AuthUtil_test.go
        ├── JsonUtil_test.go
        └── Jwt_test.go

```

- `cmd/app/main.go` : The entry point of the application. It sets up the server and initializes routes.
- `data/` : Contains JSON files that simulate data for customers, merchants, and transactions. These files are used for testing and storing history.
    
    - `customer.json` : Contains customer data (e.g., username, password, balance, is_login).
    - `merchant.json` : Contains merchant data
    - `transaction.json` : Contains transaction records.
    - Backups and sample files are also included for demonstration.

- `gin.log`: Log file for storing server logs.
- `go.mod` and `go.sum`: Go modules for dependency management.
- `internal/`: Contains the core logic and business components of the application.
    - `dto/`: Contains data transfer objects used for communication between different layers.
      - `CommonResponse.go`, `Request.go`, `Response.go`: Defines response structures and utility objects.
    - `handler/`: Contains HTTP request handlers for login and transaction functionality.
      - `Auth.go`: Handles customer login.
      - `TransactionHandler.go`: Handles payment and transaction logic.
    - `models/`: Defines data models for customer, merchant, and transaction.
      - `Customer.go,` `Merchant.go`, `Transaction.go`: Defines the structures used in the application.
    - `pkg/`: Contains utility functions for token generation and rollback.
      - `data/Rollback.go`: Handles rolling back changes in the event of errors.
      - `token/Token.go`: Responsible for creating and validating tokens.
    - `repositories/`: Contains logic for interacting with data (in this case, using JSON files).
      - `CustomerRepository.go`, `MerchantRepository.go`, `TransactionRepository.go`: Manages CRUD operations for customer, merchant, and transaction data.
    - `router/`: Contains routing logic for the application.
      - `Router.go`: Defines API routes and links them to the appropriate handlers.
    - `services/`: Contains business logic and services.
      - `AuthService.go`: Manages login, validation, and authentication logic.
      - `TransactionService.go`: Handles transaction logic, including payments.
    - `utils/`: Utility functions.
      - `AuthUtil.go`: Helper functions for handling authentication.
      - `JsonUtils.go`: Utility functions for reading and writing JSON data.
- `tests/`: Contains tests for the application.
  - `e2e/`: End-to-end tests for simulating full user interactions.
  - `integration/`: Tests for verifying integrations between components.
  - `unit/`: Unit tests for individual components like utilities and services.
  - `testdata/`: Sample JSON files used for testing.

## How To Run

---

Clone this project :

```shell
https://github.com/darulcode/merchant-bank-go-api.git
```

Install dependencies :
```shell
  go mod tidy
```

Go to data file using:
```shell
cd data
```

Check directory data address using:
```shell
pwd
```

Copy and paste on .env file.
```shell
PATH=PASTE IN HERE WHITOUT SPACE
```

run project using :
```shell
go run cmd/app/main.go
```

## API ENDPOINT

- #### Login Customer

    `POST /login`
    Request body :
    ```json
    {
        "username" : "darul",
        "password" : "123"
    }
    ```
- #### Logout Customer
    
    `POST /logout`
    
    Authorization Header using accessToken at login like:
    ```text
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoxLCJleHAiOjE3MzI1MTgwNDR9.JWsm8KeO4EfY0uMbkKlgYdfoorUiwuXxElwGA5Yhhxo
    ```

- #### Payment

    `POST /transaction`
    
    Authorization Header using accessToken at login like:
    ```text
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoxLCJleHAiOjE3MzI1MTgwNDR9.JWsm8KeO4EfY0uMbkKlgYdfoorUiwuXxElwGA5Yhhxo
    ```
    
    And Request body:
    ```json
    {
        "merchant_id": "001WPS",
        "amount": 50000
    }
    ```
    
    `GET /transaction`
    Authorization Header using accessToken at login like:
    ```text
    eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRoX2lkIjoxLCJleHAiOjE3MzI1MTgwNDR9.JWsm8KeO4EfY0uMbkKlgYdfoorUiwuXxElwGA5Yhhxo
    ```

## Testing 

---

You can run unit tests using Go’s built-in testing framework:

```shell
go test ./tests/unit
```


## Security Considerations

---

- While this API does not include advanced security mechanisms **JWT** , it ensures that only valid customers can perform actions like payments.
- For a production environment, consider implementing **token-based authentication**, **password hashing**, and **enhanced data validation**.

## Conclusion

This project provides a basic API to manage customer and merchant interactions, including login, payment, and logout. It implements logging, transaction handling, and user authentication, all while adhering to clean architecture and separation of concerns. The application is easy to set up and test, making it ideal for a simple backend solution. You can extend this project by adding advanced security features or integrating with a real database for persistent storage.






