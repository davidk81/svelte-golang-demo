CREATE TABLE "patient" (
  patientid varchar(36) not null primary key,
  name varchar(255) not null,
  location varchar(255) not null,
  created timestamp with time zone default now()
);

CREATE TABLE "user" (
  userid varchar(36) not null primary key,
  name varchar(255) not null,
  roles varchar(255) not null,
  secret varchar(1024) not null,
  created timestamp with time zone default now()
);

CREATE TABLE "patient_note" (
  noteid char(36) not null primary key,
  patientid varchar(100) not null references "patient"(patientid),
  userid varchar(100) not null references "user"(userid),
  note varchar(4096) not null,
  created timestamp with time zone default now()
);
