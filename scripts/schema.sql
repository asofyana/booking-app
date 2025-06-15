CREATE TABLE USERS (
  user_id integer primary key,
  username text not null,
  password text not null,
  full_name text not null,
  email text,
  phone_number text,
  created_date timestamp default current_timestamp not null,
  created_by text,
  updated_date timestamp default current_timestamp not null,
  updated_by text
);

CREATE TABLE ROOM (
  room_id integer primary key,
  room_name text not null,
  room_description text not null,
  room_status text not null,
  created_date timestamp default current_timestamp not null,
  created_by text,
  updated_date timestamp default current_timestamp not null,
  updated_by text
);

insert into room(room_name, room_description, room_status) values ('Room 1', 'Room 1 test', 'Active');
insert into room(room_name, room_description, room_status) values ('Room 2', 'Room 2 test', 'Active');

CREATE TABLE BOOKING (
  booking_id integer primary key,
  title text not null,
  room_id integer not null,
  booking_description text not null,
  start_date timestamp not null,
  end_date timestamp not null,
  participant_count integer,
  created_date timestamp default current_timestamp not null,
  created_by text,
  updated_date timestamp default current_timestamp not null,
  updated_by text
);