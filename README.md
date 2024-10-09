This is the public github repo solely meant for GHC 24 session "From Chaos to Clarity: Harnessing Observability in Distributed Systems".

## Overview

This is a concert ticket booking application. The user can purchase different types and quantities of tickets for the concert.


## Section-1

**Step 1: Prerequisites - Ensure Docker is installed on your machine:**

**macOS:** Use Homebrew to install Docker Desktop.

```
brew install --cask docker
```

**Windows:** Download Docker Desktop from the official Docker website.
**Linux:** Follow Docker installation instructions specific to your distribution.

**Step 2: Clone Repo**

```
git clone <repository-url>
```

**Step 3:  Bring up the docker**

```
cd <project-directory>

docker-compose up --build
```

**Step 4: In another tab, run the command to purchase the concert ticket for the concert:**

```
cd <project-directory>

curl -X POST http://localhost:8083/book-tickets \
  -H "Content-Type: application/json" \
  -d '{
        "user": {
          "id": "U123",
          "name": "John Doe",
          "dob": "1990-05-15"
        },
        "ticket": {
          "type": "VIP",
          "quantity": 2
        },
        "payment_method": {
          "type": "CreditCard",
          "number": "4111111111111111",
          "authorization": "authToken123"
        }
      }'
```


## Section-2

Navigate to the terminal where docker is running and observe the logs.

## Section-3

**Step 1: Add the missing logs - Open the following files in IDE of your choice and add the below mentioned logs**

1. ticket-registry/service/service.go

```
logEntry.WithField("ticketTypeKey", ticketTypeKey).Info("ticketTypeKey value")
```

(continue)

**Step 2: Run the following command again to purchase your concert tickets**

```
curl -X POST http://localhost:8083/book-tickets \
  -H "Content-Type: application/json" \
  -d '{
        "user": {
          "id": "U123",
          "name": "John Doe",
          "dob": "1990-05-15"
        },
        "ticket": {
          "type": "VIP",
          "quantity": 2
        },
        "payment_method": {
          "type": "CreditCard",
          "number": "4111111111111111",
          "authorization": "authToken123"
        }
      }'
```

**Step 3: Navigate to tab where docker is running & observe the logs again.**

Can you identify the issue now?


## (Optional) Section-4

**Step-1**
Open ticket-registry/mappers/tickets.go in an IDE and make the following changes to fix the code:


**Step-2**
Navigate to tab where docker is running & observe the logs again.




