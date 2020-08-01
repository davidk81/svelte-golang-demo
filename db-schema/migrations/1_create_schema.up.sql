CREATE TABLE "patient" (
  patientid varchar(255) primary key,
  name varchar(255),
  location varchar(255),
  created timestamp with time zone default now()
);

CREATE TABLE "user" (
  userid varchar(255) primary key,
  name varchar(255),
  roles varchar(255),
  secret varchar(4096),
  created timestamp with time zone default now()
);

CREATE TABLE "patient_note" (
  noteid varchar(255) primary key,
  patientid varchar(100) references "patient"(patientid),
  userid varchar(100) references "user"(userid),
  note varchar(4096),
  created timestamp with time zone default now()
);
