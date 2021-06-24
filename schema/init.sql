CREATE TABLE "User"
(
    "ID"           serial                   NOT NULL,
    "Email"        character varying        NOT NULL,
    "PasswordHash" character varying        NOT NULL,
    "Name"         character varying        NOT NULL,
    "Surname"      character varying        NOT NULL,
    "Patronymic"   character varying        NOT NULL,
    "Photo"        character varying        NOT NULL,
    "Company"      character varying        NOT NULL,
    "CreateDate"   timestamp with time zone NOT NULL,
    "Birthday"     timestamp with time zone,
    CONSTRAINT "User_pk" PRIMARY KEY ("ID")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Event"
(
    "ID"               serial                   NOT NULL,
    "Caption"          character varying        NOT NULL,
    "Description"      character varying        NOT NULL,
    "DateStart"        timestamp with time zone NOT NULL,
    "DateEnd"          timestamp with time zone NOT NULL,
    "City"             bigint                   NOT NULL,
    "Place"            character varying        NOT NULL,
    "Type"             character varying        NOT NULL,
    "Template"         character varying        NOT NULL,
    "Photo"            character varying        NOT NULL,
    "ParticipantMax"   bigint                   NOT NULL,
    "ApplicationStart" timestamp with time zone NOT NULL,
    "ApplicationEnd"   timestamp with time zone NOT NULL,
    CONSTRAINT "Event_pk" PRIMARY KEY ("ID")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Application"
(
    "User"       bigint            NOT NULL,
    "Event"      bigint            NOT NULL,
    "Role"       bigint            NOT NULL,
    "Group"      bigint            NOT NULL,
    "Data"       json              NOT NULL,
    "Organizer"  BOOLEAN           NOT NULL,
    "Identifier" character varying NOT NULL,
    CONSTRAINT "Application_pk" PRIMARY KEY ("User", "Event")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "EventRole"
(
    "Event"    bigint            NOT NULL,
    "Caption"  character varying NOT NULL,
    "Index"    serial            NOT NULL UNIQUE,
    "Checking" BOOLEAN           NOT NULL,
    CONSTRAINT "EventRole_pk" PRIMARY KEY ("Event", "Caption")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "EventGroup"
(
    "Event"   bigint            NOT NULL,
    "Caption" character varying NOT NULL,
    "Index"   serial            NOT NULL UNIQUE,
    CONSTRAINT "EventGroup_pk" PRIMARY KEY ("Event", "Caption")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Tag"
(
    "Name"  character varying NOT NULL,
    "Color" character varying NOT NULL,
    CONSTRAINT "Tag_pk" PRIMARY KEY ("Name")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "EventTag"
(
    "ID"    bigint            NOT NULL,
    "Tag"   character varying NOT NULL,
    "Event" bigint            NOT NULL,
    CONSTRAINT "EventTag_pk" PRIMARY KEY ("ID")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "City"
(
    "ID"     serial NOT NULL,
    "Region" bigint NOT NULL,
    "Name"   bigint NOT NULL,
    CONSTRAINT "City_pk" PRIMARY KEY ("ID")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Region"
(
    "ID"   serial            NOT NULL,
    "Name" character varying NOT NULL,
    CONSTRAINT "Region_pk" PRIMARY KEY ("ID")
) WITH (
      OIDS = FALSE
      );



CREATE TABLE "Log"
(
    "ID"      serial                   NOT NULL,
    "Level"   character varying        NOT NULL,
    "Message" character varying        NOT NULL,
    "Date"    timestamp with time zone NOT NULL,
    CONSTRAINT "Log_pk" PRIMARY KEY ("ID")
) WITH (
      OIDS = FALSE
      );



ALTER TABLE "Event"
    ADD CONSTRAINT "Event_fk0" FOREIGN KEY ("City") REFERENCES "City" ("ID");

ALTER TABLE "Application"
    ADD CONSTRAINT "Application_fk0" FOREIGN KEY ("User") REFERENCES "User" ("ID");
ALTER TABLE "Application"
    ADD CONSTRAINT "Application_fk1" FOREIGN KEY ("Event") REFERENCES "Event" ("ID");
ALTER TABLE "Application"
    ADD CONSTRAINT "Application_fk2" FOREIGN KEY ("Role") REFERENCES "EventRole" ("Index");
ALTER TABLE "Application"
    ADD CONSTRAINT "Application_fk3" FOREIGN KEY ("Group") REFERENCES "EventGroup" ("Index");

ALTER TABLE "EventRole"
    ADD CONSTRAINT "EventRole_fk0" FOREIGN KEY ("Event") REFERENCES "Event" ("ID");

ALTER TABLE "EventGroup"
    ADD CONSTRAINT "EventGroup_fk0" FOREIGN KEY ("Event") REFERENCES "Event" ("ID");


ALTER TABLE "EventTag"
    ADD CONSTRAINT "EventTag_fk0" FOREIGN KEY ("Tag") REFERENCES "Tag" ("Name");
ALTER TABLE "EventTag"
    ADD CONSTRAINT "EventTag_fk1" FOREIGN KEY ("Event") REFERENCES "Event" ("ID");

ALTER TABLE "City"
    ADD CONSTRAINT "City_fk0" FOREIGN KEY ("Region") REFERENCES "Region" ("ID");


