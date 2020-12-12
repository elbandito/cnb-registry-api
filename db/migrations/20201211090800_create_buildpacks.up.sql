CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS buildpacks(
   id uuid DEFAULT uuid_generate_v4 (),
   namespace VARCHAR (250) NOT NULL,
   bp_name VARCHAR (250) NOT NULL,
   version VARCHAR (250) NOT NULL,
   addr VARCHAR (250) NOT NULL,
   description TEXT NOT NULL,
   license TEXT NOT NULL,
   PRIMARY KEY (id)
);