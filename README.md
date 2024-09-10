# concert-tickets-ghc-2024

get ticket
Payment

updateTicket


Ticket:

book ticket API:

Request:
 User: Id, name , address
 Ticket: Type, quantity
 Payment: Type, Number, Authorization


Response:
    Ticket: Id, Name, Type, Quantity
   

Catalog

Get Ticket API
Request: Type, quantity
Response: []Id, Type

Update //rollback scenarios
Request: []Id, Type
Response: Status


Payment

MakePayment
Request: 
 User: Id,Name, DOB
 Payment method: Type, Number, Authorization

Response: PaymentId, Status


Ticket-Catalog
 Type
  initial Quantity
  Available


Ticket
 User:
 Transaction: Id, Type , timestamp
 Ticket Id

Transaction
 Id
 Type
 Timestamp
 Status