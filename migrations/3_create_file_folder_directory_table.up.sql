create table if not exists file.folder_directory(
      uid_parent uuid not null,
      uid_child uuid not null,
      constraint  fk_file_directory_uid_parent foreign key(uid_parent) references file.folders(uid) on update cascade on delete cascade,
      constraint  fk_file_directory_uid_child  foreign key(uid_child) references file.folders(uid)  on update cascade on delete cascade
);
