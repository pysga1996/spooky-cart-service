create table spooky_cart.cart
(
    id bigserial not null
        constraint cart_pk
            primary key,
    product_quantity integer default 0 not null,
    create_time timestamp not null,
    update_time timestamp,
    status integer not null
);

alter table spooky_cart.cart owner to postgres;

create table spooky_cart.cart_product
(
    cart_id bigserial not null
        constraint cart_product_pk
            primary key,
    product_code varchar not null,
    quantity integer not null
);

alter table spooky_cart.cart_product owner to postgres;

