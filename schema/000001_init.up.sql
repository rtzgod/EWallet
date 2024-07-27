create table wallets (
    id varchar(255) primary key,
    amount numeric not null
);

create table transactions (
    time timestamp not null,
    sender_id varchar(255) not null,
    receiver_id varchar(255) not null,
    amount numeric not null,
    foreign key(sender_id) references wallets(id) on delete cascade,
    foreign key(receiver_id) references wallets(id) on delete cascade
);