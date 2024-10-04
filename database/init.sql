
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    ticket_type VARCHAR(50) NOT NULL,
    current_quantity INT NOT NULL
);

INSERT INTO tickets (id, ticket_type, current_quantity)
VALUES (1, 'VIP', 100), (2, 'General', 200);
