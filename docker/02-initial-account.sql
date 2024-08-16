create table local_db.account
(
    account_id   int auto_increment
        primary key,
    customer_id varchar(100) not null,
    opening_date date not null,
    account_type varchar(10) not null,
    amount varchar(100) not null,
    status tinyint not null,
);

-- INSERT INTO local_db.account (id, name) VALUES (1, 'account-lion');
