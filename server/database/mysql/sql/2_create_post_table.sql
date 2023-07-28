CREATE TABLE `posts` (
  `id`          varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `title`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `github_url`        varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `app_url`     varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `slide_url`   varchar(255) COLLATE utf8mb4_bin NOT NULL,
  `score`       float        COLLATE utf8mb4_bin NOT NULL,
  `created_at`  TIMESTAMP    COLLATE utf8mb4_bin NOT NULL,
  FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE,
  FOREIGN KEY (member_id) REFERENCES members(id) ON DELETE CASCADE,
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;