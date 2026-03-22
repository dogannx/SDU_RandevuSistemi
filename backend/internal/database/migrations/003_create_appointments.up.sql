CREATE TABLE IF NOT EXISTS appointments (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    student_id  UUID NOT NULL REFERENCES students(id) ON DELETE CASCADE,
    teacher_id  UUID NOT NULL REFERENCES teachers(id) ON DELETE CASCADE,
    date        DATE NOT NULL,
    time        VARCHAR(5) NOT NULL,
    status      VARCHAR(20) DEFAULT 'upcoming',
    created_at  TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_appointments_student_id ON appointments(student_id);
CREATE INDEX IF NOT EXISTS idx_appointments_teacher_id ON appointments(teacher_id);
