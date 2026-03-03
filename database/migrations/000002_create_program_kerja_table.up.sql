CREATE TABLE program_kerja (
    internal_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    public_id CHAR(36) NOT NULL UNIQUE,
    nama_proker VARCHAR(255) NOT NULL,
    deskripsi TEXT,
    divisi VARCHAR(100),
    status VARCHAR(50),
    link_oprec VARCHAR(255),
    created_by BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_proker FOREIGN KEY (created_by) REFERENCES users(internal_id) ON DELETE SET NULL
);