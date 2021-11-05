create table agencies
(
    agencyid       integer not null
        constraint agencies_pkey
            primary key,
    name           varchar(45),
    phonenumber    varchar(45),
    specialization varchar(45)
);

alter table agencies
    owner to ydmscaoenbipqz;

create table person
(
    userid      serial
        constraint person_pkey
            primary key,
    username    varchar(45) not null,
    password    varchar(45) not null,
    firstname   varchar(45) not null,
    lastname    varchar(45) not null,
    email       varchar(45) not null,
    address     varchar(45) not null,
    phonenumber varchar(45) not null,
    role        varchar(45) not null
);

alter table person
    owner to ydmscaoenbipqz;

create table client
(
    clientid      integer not null,
    person_userid integer not null
        constraint fk_client_person1
            references person
            on update cascade on delete cascade,
    constraint client_pkey
        primary key (clientid, person_userid)
);

alter table client
    owner to ydmscaoenbipqz;

create index fk_client_person1_idx
    on client (person_userid);

create unique index client_clientid_uindex
    on client (clientid);

create table clinician
(
    clinicianid   integer not null,
    person_userid integer not null
        constraint fk_clinician_person1
            references person
            on update cascade on delete cascade,
    constraint clinician_pkey
        primary key (clinicianid, person_userid)
);

alter table clinician
    owner to ydmscaoenbipqz;

create index fk_clinician_person1_idx
    on clinician (person_userid);

create unique index clinician_clinicianid_uindex
    on clinician (clinicianid);

create table family_member
(
    familyid        serial,
    relationship    varchar(45) not null,
    client_clientid integer     not null
        constraint fk_family_member_client1
            references client (clientid)
            on update cascade on delete cascade,
    person_userid   integer     not null
        constraint fk_family_member_person1
            references person
            on update cascade on delete cascade,
    constraint family_member_pkey
        primary key (familyid, client_clientid, person_userid)
);

alter table family_member
    owner to ydmscaoenbipqz;

create index fk_family_member_client1_idx
    on family_member (client_clientid);

create index fk_family_member_person1_idx
    on family_member (person_userid);

create unique index family_member_familyid_uindex
    on family_member (familyid);

create table support_network
(
    supportid       integer not null,
    relationship    varchar(45),
    phonenumber     varchar(45),
    client_clientid integer not null
        constraint fk_support_network_client1
            references client (clientid)
            on update cascade on delete cascade,
    constraint support_network_pkey
        primary key (supportid, client_clientid)
);

alter table support_network
    owner to ydmscaoenbipqz;

create index fk_support_network_client1_idx
    on support_network (client_clientid);

create table safety_plan_has_agencies
(
    safety_plan_safetyid integer not null,
    agencies_agencyid    integer not null
        constraint fk_safety_plan_has_agencies_agencies1
            references agencies
            on update cascade on delete cascade,
    constraint safety_plan_has_agencies_pkey
        primary key (safety_plan_safetyid, agencies_agencyid)
);

alter table safety_plan_has_agencies
    owner to ydmscaoenbipqz;

create index fk_safety_plan_has_agencies_agencies1_idx
    on safety_plan_has_agencies (agencies_agencyid);

create index fk_safety_plan_has_agencies_safety_plan1_idx
    on safety_plan_has_agencies (safety_plan_safetyid);

create table client_has_clinician
(
    client_clientid       integer not null
        constraint fk_client_has_clinician_client1
            references client (clientid)
            on update cascade on delete cascade,
    clinician_clinicianid integer not null
        constraint fk_client_has_clinician_clinician1
            references clinician (clinicianid)
            on update cascade on delete cascade,
    constraint client_has_clinician_pkey
        primary key (client_clientid, clinician_clinicianid)
);

alter table client_has_clinician
    owner to ydmscaoenbipqz;

create index fk_client_has_clinician_client1_idx
    on client_has_clinician (client_clientid);

create index fk_client_has_clinician_clinician1_idx
    on client_has_clinician (clinician_clinicianid);

create table appointments
(
    appointmentid         integer not null,
    appointmenttime       timestamp,
    appointmentmedium     varchar(45),
    client_clientid       integer not null
        constraint fk_appointments_client1
            references client (clientid)
            on update cascade on delete cascade,
    clinician_clinicianid integer not null
        constraint fk_appointments_clinician1
            references clinician (clinicianid)
            on update cascade on delete cascade,
    constraint appointments_pkey
        primary key (appointmentid, client_clientid, clinician_clinicianid)
);

alter table appointments
    owner to ydmscaoenbipqz;

create index fk_appointments_client1_idx
    on appointments (client_clientid);

create index fk_appointments_clinician1_idx
    on appointments (clinician_clinicianid);

create table safety_plan
(
    safetyid              integer not null,
    triggers              varchar(500),
    warningsigns          varchar(500),
    destructivebehaviors  varchar(500),
    internalstrategies    varchar(500),
    updateddatetime       timestamp,
    updatedclinician      integer not null
        constraint fk_safety_plan_clinician2
            references clinician (clinicianid)
            on update cascade,
    client_clientid       integer not null
        constraint fk_safety_plan_client1
            references client (clientid)
            on update cascade on delete cascade,
    clinician_clinicianid integer not null,
    constraint safety_plan_pkey
        primary key (safetyid, client_clientid, clinician_clinicianid)
);

alter table safety_plan
    owner to ydmscaoenbipqz;

create index fk_safety_plan_client1_idx
    on safety_plan (client_clientid);

create index fk_safety_plan_clinician1_idx
    on safety_plan (clinician_clinicianid);

