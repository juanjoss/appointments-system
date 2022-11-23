CREATE TABLE "administrators" (
  "full_name" text NOT NULL,
  "phone_number" text UNIQUE NOT NULL,
  "email" text PRIMARY KEY UNIQUE NOT NULL,
  "password" text NOT NULL,
  "availability_period" int NOT NULL DEFAULT 1,
  "attention_hour_start" time NOT NULL DEFAULT '14:00:00',
  "attention_hour_end" time NOT NULL DEFAULT '19:00:00'
);

CREATE TABLE "appointments" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "date" date NOT NULL,
  "attention_day" text NOT NULL,
  "attention_hour" int NOT NULL,
  "status" text NOT NULL DEFAULT 'disabled',
  UNIQUE ("date", "attention_day", "attention_hour"),
  CHECK (
    status = 'disabled' OR 
    status = 'available' OR 
    status = 'reserved'
  )
);

CREATE TABLE "attention_days" (
  "day" text PRIMARY KEY NOT NULL,
  "active" boolean NOT NULL DEFAULT FALSE,
  CHECK (
    "day" = 'Lunes' OR
    "day" = 'Martes' OR
    "day" = 'Miércoles' OR
    "day" = 'Jueves' OR
    "day" = 'Viernes'
  )
);

CREATE TABLE "attention_hours" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "start_hour" time NOT NULL,
  "end_hour" time NOT NULL,
  UNIQUE ("start_hour", "end_hour"),
  CHECK (
    "end_hour" > "start_hour"
  )
);

CREATE TABLE "clients" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "full_name" text NOT NULL,
  "birthdate" date NOT NULL,
  "address" text NOT NULL,
  "locality" text NOT NULL,
  "country" text NOT NULL,
  "country_code" text NOT NULL,
  "email" text UNIQUE NOT NULL,
  "phone_number" text NOT NULL
);

CREATE TABLE "appointment_types" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" text NOT NULL,
  "description" text NOT NULL,
  "price" float NOT NULL,
  CHECK (
    price > 0
  )
);

CREATE TABLE "payment_methods" (
  "id" SERIAL PRIMARY KEY NOT NULL,
  "name" text NOT NULL
);

CREATE TABLE "bookings" (
  "client_id" int,
  "appointment_type_id" int,
  "payment_method_id" int,
  "appointment_id" int,
  PRIMARY KEY (
    "client_id", 
    "appointment_type_id", 
    "payment_method_id", 
    "appointment_id"
  )
);

ALTER TABLE "appointments" ADD FOREIGN KEY ("attention_day") REFERENCES "attention_days" ("day");
ALTER TABLE "appointments" ADD FOREIGN KEY ("attention_hour") REFERENCES "attention_hours" ("id");
ALTER TABLE "bookings" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");
ALTER TABLE "bookings" ADD FOREIGN KEY ("appointment_type_id") REFERENCES "appointment_types" ("id");
ALTER TABLE "bookings" ADD FOREIGN KEY ("payment_method_id") REFERENCES "payment_methods" ("id");
ALTER TABLE "bookings" ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id");

INSERT INTO "administrators" ("full_name", "phone_number", "email", "password") VALUES ('User', '+54 2302-123456', 'user@appointments.com', '$2a$10$wEh17KZukF8rT3haV6RlKuubAlaf2SAGATf5IoMtDcyrQX2AlbCgO');

INSERT INTO "payment_methods" ("name") VALUES ('MercadoPago');
INSERT INTO "payment_methods" ("name") VALUES ('Transferencia Bancaria');

INSERT INTO "attention_days" ("day") VALUES ('Lunes');
INSERT INTO "attention_days" ("day") VALUES ('Martes');
INSERT INTO "attention_days" ("day") VALUES ('Miércoles');
INSERT INTO "attention_days" ("day") VALUES ('Jueves');
INSERT INTO "attention_days" ("day") VALUES ('Viernes');

INSERT INTO "attention_hours" ("start_hour", "end_hour") VALUES ('14:00:00', '15:00:00');
INSERT INTO "attention_hours" ("start_hour", "end_hour") VALUES ('15:00:00', '16:00:00');
INSERT INTO "attention_hours" ("start_hour", "end_hour") VALUES ('16:00:00', '17:00:00');
INSERT INTO "attention_hours" ("start_hour", "end_hour") VALUES ('17:00:00', '18:00:00');
INSERT INTO "attention_hours" ("start_hour", "end_hour") VALUES ('18:00:00', '19:00:00');