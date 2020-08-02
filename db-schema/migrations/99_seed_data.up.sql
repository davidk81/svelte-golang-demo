INSERT INTO "patient" (patientid, name, location) 
    VALUES ('patient1', 'patrick', 'room101');

INSERT INTO "patient" (patientid, name, location) 
    VALUES ('patient2', 'peter', 'room102');

INSERT INTO "patient" (patientid, name, location) 
    VALUES ('patient3', 'paul', 'room103');

INSERT INTO "user" (userid, name, roles, secret) 
    VALUES ('nurse1', 'nancy', 'nurse', '$2a$04$P4ouhozPJZX8NCCm7QyrIe1ZR46HNKL5tZgr0Yn4RCPyY85hnAM0m');

INSERT INTO "user" (userid, name, roles, secret) 
    VALUES ('admin', 'Administrator', 'admin', '$2a$04$P4ouhozPJZX8NCCm7QyrIe1ZR46HNKL5tZgr0Yn4RCPyY85hnAM0m');
