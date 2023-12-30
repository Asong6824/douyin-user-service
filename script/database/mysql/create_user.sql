DROP TABLE IF EXISTS user;
CREATE TABLE IF NOT EXISTS user(
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  password VARCHAR(255),
  follow_count BIGINT DEFAULT 0,       
  follower_count BIGINT DEFAULT 0,      
  avatar VARCHAR(255),
  background_image VARCHAR(255),
  signature VARCHAR(255),
  total_favorited BIGINT DEFAULT 0,     
  work_count BIGINT DEFAULT 0,          
  favorite_count BIGINT DEFAULT 0       
);

