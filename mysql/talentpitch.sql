DROP TABLE IF EXISTS program_participants;
DROP TABLE IF EXISTS challenges;
DROP TABLE IF EXISTS companies;
DROP TABLE IF EXISTS programs;
DROP TABLE IF EXISTS users;

CREATE TABLE users (
   id INT PRIMARY KEY,
   name VARCHAR(255),
   email VARCHAR(255) UNIQUE,
   image_path VARCHAR(255) NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE challenges (
    id INT PRIMARY KEY,
    title VARCHAR(255),
    description TEXT,
    difficulty INT,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE companies (
   id INT PRIMARY KEY,
   name VARCHAR(255),
   image_path VARCHAR(255) NULL,
   location VARCHAR(255),
   industry VARCHAR(255),
   user_id INT,
   FOREIGN KEY (user_id) REFERENCES users(id),
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE programs (
  id INT PRIMARY KEY,
  title VARCHAR(255),
  description TEXT,
  start_date DATE,
  end_date DATE,
  user_id INT,
  FOREIGN KEY (user_id) REFERENCES users(id),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE program_participants (
  id INT PRIMARY KEY,
  program_id INT,
  challenge_id INT,
  company_id INT,
  user_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (program_id) REFERENCES programs(id),
  FOREIGN KEY (challenge_id) REFERENCES challenges(id),
  FOREIGN KEY (company_id) REFERENCES companies(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);

alter table users
    modify id int auto_increment;

alter table users
    auto_increment = 1;

alter table challenges
    modify id int auto_increment;

alter table challenges
    auto_increment = 1;

alter table companies
    modify id int auto_increment;

alter table companies
    auto_increment = 1;

alter table programs
    modify id int auto_increment;

alter table programs
    auto_increment = 1;

alter table program_participants
    modify id int auto_increment;

alter table program_participants
    auto_increment = 1;