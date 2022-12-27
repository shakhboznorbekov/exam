   CREATE TABLE category (
        category_id UUID NOT NULL UNIQUE,
        parent_id UUID,
        category_name VARCHAR NOT NULL, 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    ); 

    CREATE TABLE product (
        product_id UUID NOT NULL PRIMARY KEY,
        product_name VARCHAR NOT NULL,
        price  INTEGER NOT NULL,
        category_id UUID NOT NULL REFERENCES category(category_id), 
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, 
        updated_at TIMESTAMP NOT NULL
    );


    -- CREATE TABLE order (
    --     id UUID NOT NULL PRIMARY KEY,
    --     description VARCHAR(255),
    --     product product 
    -- )



    -- insert into category(category_id,parent_id,category_name,updated_at)VALUES
    -- ('f23fa31a-8091-11ed-a1eb-0242ac120002',NULL,'Avtomobil',now()),
    -- ('d8d56125-ce1e-4f9d-bd38-bceab0d5f192',NULL,'Telefon',now());
    -- insert into product(product_id,product_name,price,category_id,updated_at)VALUES
    -- ('8b85ff2a-8091-11ed-a1eb-0242ac120002', 'Radar',125000,'f23fa31a-8091-11ed-a1eb-0242ac120002',now()),
    -- ('28050fd6-ea96-4db3-83d1-ba3ccd882297', 'Magicar',200000,'f23fa31a-8091-11ed-a1eb-0242ac120002',now()),
    -- ('e3113cf4-ba91-4cbf-a7fe-7ec2146ddbca', 'Redmi',500000,'d8d56125-ce1e-4f9d-bd38-bceab0d5f192',now()),
    -- ('4aab5875-b533-4cdb-9ff5-999cbd4dbab4', 'Samsung',700000,'d8d56125-ce1e-4f9d-bd38-bceab0d5f192',now());
    

    --  insert into category(category_id,parent_id,category_name,updated_at)VALUES
    -- ('fe78ce65-a4f8-41fe-960c-971c63349cef','dd8ae488-bd46-472a-a6c0-2a281da8394c','Moshina',now()),
    -- ('fc5d8b04-ecfd-4f05-9ac9-74c3c832d167','a26883c1-4607-4f24-9701-6e0d91ef3a8d','Vertalyot',now());

-- INSERT INTO users(user_id, first_name, last_name, phone_number, updated_at) VALUES
-- ('8b85ff2a-8091-11ed-a1eb-0242ac120002','Samandar', 'Foziljonov', '997191323', now()); 

-- INSERT INTO books(book_id, title, author, price, updated_at) VALUES
-- ('f23fa31a-8091-11ed-a1eb-0242ac120002', 'Tokaki', 'Murakami', 1500, now());

-- INSERT INTO orders(order_id, user_id, book_id, payed) VALUES
-- ('5af01b9c-8092-11ed-a1eb-0242ac120002', '8b85ff2a-8091-11ed-a1eb-0242ac120002', 'f23fa31a-8091-11ed-a1eb-0242ac120002', 1500);

-- SELECT 
--     users.first_name || ' ' || users.last_name as fullname,
--     SUM(orders.payed)
-- FROM
--     orders
-- JOIN users ON orders.user_id = users.user_id
-- GROUP BY fullname;