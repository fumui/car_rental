CREATE TABLE cars(
    id SERIAL PRIMARY KEY NOT NULL,
    name CHAR(50) NOT NULL,
    day_rate DECIMAL(10, 2) NOT NULL,
    month_rate DECIMAL(10, 2) NOT NULL,
    image CHAR(256) NOT NULL
);

CREATE TABLE orders(
    id SERIAL PRIMARY KEY NOT NULL,
    car_id INT NOT NULL REFERENCES cars(id),
    order_date DATE NOT NULL,
    pickup_date DATE NOT NULL,
    dropoff_date DATE NOT NULL,
    pickup_location CHAR(50) NOT NULL,
    dropoff_location CHAR(50) NOT NULL
);

