CREATE TABLE tours (
    id SERIAL PRIMARY KEY,
    chillplace TEXT NOT NULL,
    fromplace TEXT NOT NULL,
    date TIMESTAMP,
    duration INT CHECK (duration > 0),
    cost INT CHECK (cost > 0),
    tourists_number INT CHECK (tourists_number > 0),
    chilltype TEXT CHECK (chilltype in ('элитный', 'городской', 'пляжный', 'активный', 'необычный'))
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    login TEXT NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE sales (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    expiredtime TIMESTAMP,
    percent INT CHECK (percent > 0 AND percent <= 100)
);

CREATE TABLE clients (
    id SERIAL PRIMARY KEY,
    acc_id INT REFERENCES accounts(id),
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    mail VARCHAR(50) CHECK (mail SIMILAR TO '[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}'),
    phone VARCHAR(11) CHECK (phone ~ '^\d{11}$')
);

CREATE TABLE managers (
    id SERIAL PRIMARY KEY,
    acc_id INT REFERENCES accounts(id),
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    department TEXT NOT NULL
);

CREATE TABLE requests (
    id SERIAL PRIMARY KEY,
    tour_id INT REFERENCES tours(id),
    clnt_id INT REFERENCES clients(id),
    mngr_id INT REFERENCES managers(id),
    status TEXT CHECK (status in ('принята', 'обрабатывается', 'подтверждена', 'отклонена', 'оплачена')),
    create_time TIMESTAMP,
    modify_time TIMESTAMP,
    data JSON
);

CREATE TABLE tours_sales (
    tour_id INT REFERENCES tours(id),
    sale_id INT REFERENCES sales(id)
);

CREATE TABLE request_outbox (
    id SERIAL PRIMARY KEY,
    req_id INT REFERENCES requests(id),
    sum INT CHECK (sum > 0),
    state TEXT CHECK (state in ('non-send', 'send'))
);
