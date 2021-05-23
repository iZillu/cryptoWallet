CREATE TABLE users
(
    id           serial       not null unique,
    name         varchar(255) not null,
    surname      varchar(255) not null,
    email        varchar(255) not null unique ,
    password     varchar(255) not null unique,
    ip           varchar(255) not null unique,
    role         varchar(255) ,
    userAgent    varchar(255) not null,
    isVerified   boolean not null default false
);

CREATE TABLE wallets
(
    id          serial       not null unique,
    address     varchar(255) not null unique,
    balance     float
);

CREATE TABLE users_wallets
(
    id          serial                                           not null unique,
    user_id     int references users (id) on delete cascade      not null,
    wallet_id   int references wallets (id) on delete cascade    not null
);

CREATE TABLE transactions
(
    id          serial       not null unique,
    amount varchar(255) not null,
    commission varchar(255) not null,
    coinName varchar(255) not null,
    senderAddress varchar(255) not null,
    receiverAddress varchar(255) not null,
    actionTime timestamp not null  default now(),
    unixActionTime varchar(255) not null
);


CREATE TABLE users_transactions
(
    id              serial                                              not null unique,
    user_id         int references users (id) on delete cascade         not null,
    transactions_id int references transactions (id) on delete cascade  not null
);