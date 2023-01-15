create schema splash_inventory;

create table splash_inventory.cart
(
    id bigserial not null
        constraint cart_pk
            primary key,
    product_quantity integer default 0 not null,
    create_time timestamp not null,
    update_time timestamp,
    status integer not null
);

alter table splash_inventory.cart owner to postgres;

create table splash_inventory.cart_product
(
    cart_id bigserial not null
        constraint cart_product_pk
            primary key,
    product_code varchar not null,
    quantity integer not null
);

alter table splash_inventory.cart_product owner to postgres;

