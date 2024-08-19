create table local_db.accounts
(
    account_id   int auto_increment
        primary key,
    customer_id int not null,
    opening_date datetime not null,
    account_type varchar(10) not null,
    amount decimal(10, 2) not null,
    status tinyint not null,
);

INSERT INTO local_db.accounts (account_id, customer_id, opening_date, account_type, amount, status) VALUES (1, 1, '2024-08-19 10:52:16', 'AA', 100.00, 10);
INSERT INTO local_db.accounts (account_id, customer_id, opening_date, account_type, amount, status) VALUES (2, 1, '2024-08-18 11:53:58', 'BB', 500.00, 10);
