create table if not exists deliveries (
    id serial not null,
    truck_id int not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_deliveries primary key(id)
);

create table if not exists payments(
    id serial not null,
    name text not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_payments primary key(id),
    unique (name)
);

create table if not exists invoice_lines (
    id serial not null,
    item_id int not null,
    invoice_id int not null,
    quantity numeric(18,3) not null,
    amount numeric(18,3)  null,
    total_amount numeric(18,3)  not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_invoice_lines primary key(id)
);

create table if not exists invoices (
    id serial not null,
    payment_id int not null,
    client_id int null,
    total_amount numeric(18,3)  not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_invoices primary key(id)
);

create table if not exists clients (
    id serial not null,
    account_id int,
    name text not null,
    lastname text not null,
    primary_phone text,
    second_phone text,
    address text,
    email text,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_clients primary key(id)
);

create table if not exists accounts (
    id serial not null,
    actual_amount numeric(18,3)  not null,
    previous_amount numeric(18,3) ,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_accounts primary key(id)
);

create table if not exists roles (
    id serial not null,
    name text not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_roles primary key(id),
    unique (name)
);

create table if not exists employees (
    id serial not null,
    role_id int not null,
    name text not null,
    lastname text not null,
    phone text,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_employees primary key(id),
    unique (name, lastname)
);



CREATE TABLE IF NOT EXISTS trucks (
    id serial not null,
    employee_id int not null,
    name text not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_trucks primary key(id),
    unique (name)
);

CREATE TABLE IF NOT EXISTS warehouses (
    id serial not null,
    name text not null,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    constraint pk_warehouses primary key(id),
    unique (name)
);

CREATE TABLE IF NOT EXISTS items (
    id serial NOT NULL,
    warehouse_id int not null,
    delivery_id int,
    name text NOT NULL,
    description text NOT NULL,
    price numeric(18,3) NOT NULL,
    created_at timestamp DEFAULT now(),
    updated_at timestamp NOT NULL,
    CONSTRAINT pk_items PRIMARY KEY(id)
);

alter table items add CONSTRAINT fk_items_warehouses FOREIGN KEY(warehouse_id) REFERENCES warehouses(id);
alter table items add CONSTRAINT fk_items_deliveries FOREIGN KEY(delivery_id) REFERENCES deliveries(id);

alter table trucks add constraint fk_trucks_employees foreign key(employee_id) references employees(id);

alter table employees add constraint fk_employees_roles foreign key(role_id) references roles(id);

alter table invoice_lines add constraint fk_invoices_lines_items foreign key(item_id) references items(id);
alter table invoice_lines add constraint fk_invoices_lines_invoices foreign key(invoice_id) references invoices(id);

alter table invoices add constraint fk_invoices_payments foreign key(payment_id) references payments(id);
alter table invoices add constraint fk_invoices_clients foreign key(client_id) references clients(id);

alter table clients add constraint fk_clients_accounts foreign key(account_id) references accounts(id);