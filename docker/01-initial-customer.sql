create table local_db.customers (
    customer_id int auto_increment primary key,
    name varchar(100) not null,
    date_of_birth date not null,
    city varchar(100) not null,
    zipcode varchar(10) not null,
    status tinyint not null
);

INSERT INTO local_db.customers (customer_id, name, date_of_birth, city, zipcode, status) VALUES (1, 'Mr.Somchai', '2024-08-15', 'Bangkok', '11111', 10);
