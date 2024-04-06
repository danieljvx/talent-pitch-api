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
  entity_type VARCHAR(50), -- Indicar el tipo de entidad: 'user', 'challenge', 'company'
  entity_id INT, -- ID de la entidad participante
  FOREIGN KEY (program_id) REFERENCES programs(id),
  FOREIGN KEY (entity_id, entity_type) REFERENCES users(id, 'user') -- Ajustar para otras entidades
);