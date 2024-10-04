
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    ticket_type VARCHAR(50) NOT NULL,
    current_quantity INT NOT NULL
);

INSERT INTO tickets (id, ticket_type, current_quantity)
VALUES (1, 'VIPFrontRow', 1000), (2, 'PlatinumSeating', 2000), (3, 'GeneralAdmission', 5000), (4, 'BalconySeating', 1500), (5, 'SuperfanPit', 500);
