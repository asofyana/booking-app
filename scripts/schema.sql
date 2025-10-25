CREATE TABLE USERS (
  user_id integer primary key,
  username text not null,
  password text not null,
  full_name text not null,
  title text,
  email text,
  phone_number text,
  role text,
  status text not null,
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
  css_class text,
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
  start_date timestamp not null,
  end_date timestamp not null,
  participant_count integer,
  activity_code text,
  organizer text,
  pic text,
  pic_contactno text,
  reject_reason text,
  status text,
  created_date timestamp default current_timestamp not null,
  created_by text,
  updated_date timestamp default current_timestamp not null,
  updated_by text
);

CREATE TABLE BOOKING_ROOM (
  booking_id integer not null,
  room_id integer not null
)

create table LOOKUP (
  lookup_type text not null,
  lookup_code text not null,
  lookup_text text not null,
  lookup_status text not null
)

insert into lookup(lookup_type, lookup_code, lookup_text) values ('ACTIVITY', 'ACT_1', 'Activity 1');
insert into lookup(lookup_type, lookup_code, lookup_text) values ('ACTIVITY', 'ACT_2', 'Activity 2');

insert into lookup(lookup_type, lookup_code, lookup_text) values ('USER_ROLE', 'admin', 'admin');
insert into lookup(lookup_type, lookup_code, lookup_text) values ('USER_ROLE', 'user', 'user');
insert into lookup(lookup_type, lookup_code, lookup_text) values ('USER_STATUS', 'active', 'active');
insert into lookup(lookup_type, lookup_code, lookup_text) values ('USER_STATUS', 'inactive', 'inactive');
