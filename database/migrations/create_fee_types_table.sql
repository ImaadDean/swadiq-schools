-- Migration: Create fee_types table
-- This migration creates the fee_types table and populates it with default data

CREATE TABLE IF NOT EXISTS fee_types (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    code VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    is_active BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_fee_types_code ON fee_types(code);
CREATE INDEX IF NOT EXISTS idx_fee_types_is_active ON fee_types(is_active);
CREATE INDEX IF NOT EXISTS idx_fee_types_deleted_at ON fee_types(deleted_at);

-- Insert default fee types
INSERT INTO fee_types (id, name, code, description, is_active) VALUES
    (gen_random_uuid(), 'Tuition Fee', 'tuition', 'Regular tuition fees for academic instruction', true),
    (gen_random_uuid(), 'Examination Fee', 'examination', 'Fees for examinations and assessments', true),
    (gen_random_uuid(), 'Library Fee', 'library', 'Fees for library services and resources', true),
    (gen_random_uuid(), 'Laboratory Fee', 'laboratory', 'Fees for laboratory usage and materials', true),
    (gen_random_uuid(), 'Sports Fee', 'sports', 'Fees for sports activities and facilities', true),
    (gen_random_uuid(), 'Transport Fee', 'transport', 'Fees for school transportation services', true),
    (gen_random_uuid(), 'Other Fee', 'other', 'Miscellaneous fees not covered by other categories', true)
ON CONFLICT (code) DO NOTHING;

-- Update the fees table to add fee_type_id column if it doesn't exist
-- Note: You may need to adjust this based on your existing fees table structure
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.columns 
                   WHERE table_name = 'fees' AND column_name = 'fee_type_id') THEN
        ALTER TABLE fees ADD COLUMN fee_type_id UUID;
        
        -- Add foreign key constraint
        ALTER TABLE fees ADD CONSTRAINT fk_fees_fee_type_id 
            FOREIGN KEY (fee_type_id) REFERENCES fee_types(id);
        
        -- Create index for the foreign key
        CREATE INDEX idx_fees_fee_type_id ON fees(fee_type_id);
    END IF;
END $$;
