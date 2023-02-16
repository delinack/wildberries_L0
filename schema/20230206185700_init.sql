-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS orders
(
    order_uid          text      not null PRIMARY KEY,
    track_number       text      not null,
    entry              text      not null,
    locale             text      not null,
    internal_signature text      not null,
    customer_id        text      not null,
    delivery_service   text      not null,
    shardkey           text      not null,
    sm_id              int       not null,
    date_created       timestamp not null,
    oof_shard          text      not null
);

CREATE TABLE IF NOT EXISTS payments
(
    payment_id    uuid not null PRIMARY KEY unique default uuid_generate_v4(),
    transaction   text not null,
    request_id    text not null,
    currency      text not null,
    provider      text not null,
    amount        int  not null,
    payment_dt    int  not null,
    bank          text not null,
    delivery_cost int  not null,
    goods_total   int  not null,
    custom_fee    int  not null
);

CREATE TABLE IF NOT EXISTS deliveries
(
    delivery_id uuid not null unique PRIMARY KEY default uuid_generate_v4(),
    name        text not null,
    phone       text not null,
    zip         text not null,
    city        text not null,
    address     text not null,
    region      text not null,
    email       text not null
);

CREATE TABLE IF NOT EXISTS items
(
    item_id      uuid not null PRIMARY KEY unique default uuid_generate_v4(),
    chrt_id      int  not null,
    track_number text not null,
    price        int  not null,
    rid          text not null,
    name         text not null,
    sale         int  not null,
    size         text not null,
    total_price  int  not null,
    nm_id        int  not null,
    brand        text not null,
    status       int  not null
);

CREATE TABLE IF NOT EXISTS orders_items
(
    id        uuid                                                 not null primary key unique default uuid_generate_v4(),
    order_uid text references orders (order_uid) on delete cascade not null,
    item_id   uuid references items (item_id) on delete cascade    not null
);

CREATE TABLE IF NOT EXISTS orders_delivery
(
    id          uuid                                                       not null primary key unique default uuid_generate_v4(),
    order_uid   text references orders (order_uid) on delete cascade       not null,
    delivery_id uuid references deliveries (delivery_id) on delete cascade not null
);

CREATE TABLE IF NOT EXISTS orders_payment
(
    id         uuid                                                    not null primary key unique default uuid_generate_v4(),
    order_uid  text references orders (order_uid) on delete cascade    not null,
    payment_id uuid references payments (payment_id) on delete cascade not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS payments;
DROP TABLE IF EXISTS deliveries;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS orders_items;
DROP TABLE IF EXISTS orders_delivery;
DROP TABLE IF EXISTS orders_payment;
-- +goose StatementEnd
