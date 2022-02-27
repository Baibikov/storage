create table if not exists file.files (
    uid uuid primary key not null default gen_random_uuid(),
    file_name varchar(255) not null,
    folder uuid not null,
    added_at  timestamptz  not null default now(),
    constraint  fk_file_files_folder_uid foreign key(folder) references file.folders(uid) on delete cascade
);