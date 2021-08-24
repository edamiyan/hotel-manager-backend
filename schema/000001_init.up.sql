CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE rooms
(
    id          serial not null unique,
    room_number int    not null,
    double_bed  int    not null,
    single_bed  int    not null,
    description varchar(255),
    price       int    not null
);

CREATE TABLE bookings
(
    id             serial       not null unique,
    name           varchar(255) not null,
    phone          varchar(255) not null,
    arrival_date   date         not null,
    departure_date date         not null,
    guests_number  int          not null,
    is_booking     boolean      not null default false,
    comment        varchar(255),
    status         int          not null default 0
);

CREATE TABLE users_rooms
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    room_id int references rooms (id) on delete cascade not null
);

CREATE TABLE rooms_bookings
(
    id      serial                                         not null unique,
    item_id int references bookings (id) on delete cascade not null,
    list_id int references rooms (id) on delete cascade    not null
);