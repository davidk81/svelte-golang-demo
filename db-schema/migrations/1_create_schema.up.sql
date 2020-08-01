CREATE TABLE "patient" (
  patientid varchar(255) not null primary key,
  name varchar(255) not null,
  location varchar(255) not null,
  created timestamp with time zone default now()
);

CREATE TABLE "user" (
  userid varchar(255) not null primary key,
  name varchar(255) not null,
  roles varchar(255) not null,
  secret varchar(4096),
  created timestamp with time zone default now()
);

CREATE TABLE "patient_note" (
  noteid char(64) not null primary key,
  patientid varchar(100) not null references "patient"(patientid),
  userid varchar(100) not null references "user"(userid),
  note varchar(4096) not null,
  created timestamp with time zone default now()
);
