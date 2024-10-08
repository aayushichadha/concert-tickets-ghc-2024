This is the public github repo solely meant for GHC 24 session "From Chaos to Clarity: Harnessing Observability in Distributed Systems".

**Overview**

This is a concert ticket booking application. The user can purchase different types and quantities of tickets for the concert.


**Section-1**

Run the following command to purchase the concert ticket for the Taylor Swift concert:

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



