
CREATE TABLE IF NOT EXISTS tickets (
    id SERIAL PRIMARY KEY,
    ticket_type VARCHAR(50) NOT NULL,
    current_quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS ticket_bookings (
    id VARCHAR(255) PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    type VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL
    );


INSERT INTO tickets (id, ticket_type, current_quantity)
VALUES (1, 'VIPFrontRow', 1000), (2, 'PlatinumSeats', 2000), (3, 'GeneralAdmission', 5000), (4, 'BalconySeat', 1500), (5, 'SuperfanPit', 500);
