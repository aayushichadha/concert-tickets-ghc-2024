This is the public github repo solely meant for GHC 24 session "From Chaos to Clarity: Harnessing Observability in Distributed Systems".

## Overview

This is a concert ticket booking application. The user can purchase different types and quantities of tickets for the concert.


## Section-1

**Step 1: Prerequisites - Ensure Docker is installed on your machine:**

Visit official ([Docker website](https://www.docker.com/products/docker-desktop/)) to download Docker or use the commands below.

**macOS:** Use Homebrew to install Docker Desktop.

```
brew install --cask docker
```

**Linux:** Follow Docker installation instructions specific to your distribution.

```
sudo apt update
sudo apt install docker-ce docker-ce-cli containerd.io
```

**Step 2: Clone Repo**

```
git clone https://github.com/aayushichadha/concert-tickets-ghc-2024.git
```

**Step 3:  Bring up the docker**

```
cd concert-tickets-ghc-2024

docker-compose up --build
```

**Step 4: In another tab, run the command to purchase the concert ticket for the concert:**

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
          "type": "general-admission",
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

**Step 1: Add the missing logs - Go through the code and add more logs to debug and find out the issue**

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
          "type": "general-admission",
          "quantity": 2
        },
        "payment_method": {
          "type": "CreditCard",
          "number": "4111111111111111",
          "authorization": "authToken123"
        }
      }'
```

**Step 3: Navigate to tab where docker is running & observe the enhanced logs.**

Can you identify the issue now?


## (Optional) Section-4

**Step-1**
Make the changes to fix the code and restart the docker by following commands:

```
docker-compose down
docker-compose up --build
```


**Step-2**
Navigate to tab where docker is running & observe the logs again.


