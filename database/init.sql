
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    ticket_type VARCHAR(50) NOT NULL,
    current_quantity INT NOT NULL
);

INSERT INTO tickets (ticket_type, current_quantity)
VALUES ('VIP', 100), ('General', 200);
