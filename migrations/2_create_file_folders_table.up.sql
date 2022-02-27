create table if not exists file.folders (
  uid uuid primary key not null default gen_random_uuid(),
  name varchar(255) not null,
  level integer not null default 0,
  created_at timestamptz not null default now()
);