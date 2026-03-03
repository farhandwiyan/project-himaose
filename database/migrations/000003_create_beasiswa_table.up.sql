CREATE TABLE beasiswa (
    internal_id BIGINT PRIMARY KEY AUTO_INCREMENT,
    public_id CHAR(36) NOT NULL UNIQUE,
    nama_beasiswa VARCHAR(255) NOT NULL,
    link_pendaftaran VARCHAR(255),
    tgl_buka DATE,
    tgl_tutup DATE,
    created_by BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_beasiswa FOREIGN KEY (created_by) REFERENCES users(internal_id) ON DELETE SET NULL
);