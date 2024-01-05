create extension if not exists pg_trgm;

create table services (
  id uuid primary key,
  sequence bigserial,
  name text not null,
  description text not null,
  created_at timestamp without time zone not null default now()
);

create index trgm_services_name on services using gin (name gin_trgm_ops);
create index trgm_services_description on services using gin (to_tsvector('english', description));

create table versions (
  id uuid primary key,
  service_id uuid not null,
  version text not null,
  created_at timestamp without time zone not null default now(),
  foreign key (service_id) references services (id)
);
