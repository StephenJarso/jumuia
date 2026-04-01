-- Mock Data for Jumuia VSLA Platform
-- Realistic data for testing and demonstration

-- Insert Seasons
INSERT INTO seasons (name, year) VALUES 
('Long Rains', 2025),
('Short Rains', 2025),
('Long Rains', 2026);

-- Insert Crops
INSERT INTO crops (name) VALUES 
('Maize'),
('Beans'),
('Sorghum'),
('Millet'),
('Groundnuts'),
('Cassava'),
('Sweet Potatoes'),
('Vegetables');

-- Insert Disasters
INSERT INTO disasters (type, description, location, date) VALUES 
('Drought', 'Severe drought affecting crop production in the region', 'Kisumu County', '2025-06-15'),
('Flood', 'Heavy rainfall causing flooding in low-lying areas', 'Busia County', '2025-09-20'),
('Pest Invasion', 'Army worm invasion destroying maize crops', 'Siaya County', '2025-07-10');

-- Insert Groups (without leader_id initially)
INSERT INTO groups (name, village, district) VALUES 
('Tumaini VSLA', 'Kisumu Village', 'Kisumu Central'),
('Upendo Group', 'Busia Town', 'Busia County'),
('Amani Savings', 'Siaya Center', 'Siaya County'),
('Umoja Women', 'Mumias Village', 'Kakamega County'),
('Vijana Youth', 'Bungoma Town', 'Bungoma County');

-- Insert Members for Group 1: Tumaini VSLA
INSERT INTO members (group_id, name, phone, role) VALUES 
(1, 'Mary Akinyi', '+254712345678', 'leader'),
(1, 'John Odhiambo', '+254723456789', 'treasurer'),
(1, 'Grace Wanjiku', '+254734567890', 'member'),
(1, 'Peter Kamau', '+254745678901', 'member'),
(1, 'Alice Muthoni', '+254756789012', 'member'),
(1, 'David Ochieng', '+254767890123', 'member'),
(1, 'Sarah Njeri', '+254778901234', 'member');

-- Insert Members for Group 2: Upendo Group
INSERT INTO members (group_id, name, phone, role) VALUES 
(2, 'James Barasa', '+254789012345', 'leader'),
(2, 'Elizabeth Nafula', '+254790123456', 'treasurer'),
(2, 'Michael Wekesa', '+254701234567', 'member'),
(2, 'Catherine Nasimiyu', '+254712345670', 'member'),
(2, 'Robert Wanyama', '+254723456781', 'member');

-- Insert Members for Group 3: Amani Savings
INSERT INTO members (group_id, name, phone, role) VALUES 
(3, 'Joseph Otieno', '+254734567892', 'leader'),
(3, 'Florence Adhiambo', '+254745678903', 'treasurer'),
(3, 'Daniel Omondi', '+254756789014', 'member'),
(3, 'Mercy Akoth', '+254767890125', 'member'),
(3, 'Samuel Oduor', '+254778901236', 'member'),
(3, 'Rose Atieno', '+254789012347', 'member');

-- Insert Members for Group 4: Umoja Women
INSERT INTO members (group_id, name, phone, role) VALUES 
(4, 'Agnes Wambui', '+254790123458', 'leader'),
(4, 'Lucy Mwikali', '+254701234569', 'treasurer'),
(4, 'Hannah Nduku', '+254712345671', 'member'),
(4, 'Esther Mumbua', '+254723456782', 'member'),
(4, 'Ruth Kavutha', '+254734567893', 'member'),
(4, 'Deborah Mwende', '+254745678904', 'member'),
(4, 'Miriam Syombua', '+254756789015', 'member'),
(4, 'Tabitha Kanini', '+254767890126', 'member');

-- Insert Members for Group 5: Vijana Youth
INSERT INTO members (group_id, name, phone, role) VALUES 
(5, 'Kevin Ouma', '+254778901237', 'leader'),
(5, 'Brian Kimani', '+254789012348', 'treasurer'),
(5, 'Dennis Mutua', '+254790123459', 'member'),
(5, 'Felix Kipchoge', '+254701234570', 'member'),
(5, 'George Mwangi', '+254712345672', 'member'),
(5, 'Henry Njoroge', '+254723456783', 'member');

-- Update Groups with leader_id (now that members exist)
UPDATE groups SET leader_id = 1 WHERE id = 1; -- Mary Akinyi
UPDATE groups SET leader_id = 8 WHERE id = 2; -- James Barasa
UPDATE groups SET leader_id = 13 WHERE id = 3; -- Joseph Otieno
UPDATE groups SET leader_id = 19 WHERE id = 4; -- Agnes Wambui
UPDATE groups SET leader_id = 27 WHERE id = 5; -- Kevin Ouma

-- Insert Savings (realistic amounts in KES)
INSERT INTO savings (member_id, amount, meeting_date) VALUES 
-- Group 1: Tumaini VSLA
(1, 500.00, '2025-01-15'),
(1, 500.00, '2025-02-15'),
(1, 500.00, '2025-03-15'),
(2, 500.00, '2025-01-15'),
(2, 500.00, '2025-02-15'),
(2, 500.00, '2025-03-15'),
(3, 300.00, '2025-01-15'),
(3, 300.00, '2025-02-15'),
(3, 300.00, '2025-03-15'),
(4, 400.00, '2025-01-15'),
(4, 400.00, '2025-02-15'),
(4, 400.00, '2025-03-15'),
(5, 350.00, '2025-01-15'),
(5, 350.00, '2025-02-15'),
(5, 350.00, '2025-03-15'),
(6, 450.00, '2025-01-15'),
(6, 450.00, '2025-02-15'),
(6, 450.00, '2025-03-15'),
(7, 300.00, '2025-01-15'),
(7, 300.00, '2025-02-15'),
(7, 300.00, '2025-03-15'),

-- Group 2: Upendo Group
(8, 600.00, '2025-01-20'),
(8, 600.00, '2025-02-20'),
(8, 600.00, '2025-03-20'),
(9, 600.00, '2025-01-20'),
(9, 600.00, '2025-02-20'),
(9, 600.00, '2025-03-20'),
(10, 400.00, '2025-01-20'),
(10, 400.00, '2025-02-20'),
(10, 400.00, '2025-03-20'),
(11, 500.00, '2025-01-20'),
(11, 500.00, '2025-02-20'),
(11, 500.00, '2025-03-20'),
(12, 450.00, '2025-01-20'),
(12, 450.00, '2025-02-20'),
(12, 450.00, '2025-03-20'),

-- Group 3: Amani Savings
(13, 800.00, '2025-01-10'),
(13, 800.00, '2025-02-10'),
(13, 800.00, '2025-03-10'),
(14, 800.00, '2025-01-10'),
(14, 800.00, '2025-02-10'),
(14, 800.00, '2025-03-10'),
(15, 600.00, '2025-01-10'),
(15, 600.00, '2025-02-10'),
(15, 600.00, '2025-03-10'),
(16, 700.00, '2025-01-10'),
(16, 700.00, '2025-02-10'),
(16, 700.00, '2025-03-10'),
(17, 650.00, '2025-01-10'),
(17, 650.00, '2025-02-10'),
(17, 650.00, '2025-03-10'),
(18, 750.00, '2025-01-10'),
(18, 750.00, '2025-02-10'),
(18, 750.00, '2025-03-10'),

-- Group 4: Umoja Women
(19, 1000.00, '2025-01-25'),
(19, 1000.00, '2025-02-25'),
(19, 1000.00, '2025-03-25'),
(20, 1000.00, '2025-01-25'),
(20, 1000.00, '2025-02-25'),
(20, 1000.00, '2025-03-25'),
(21, 800.00, '2025-01-25'),
(21, 800.00, '2025-02-25'),
(21, 800.00, '2025-03-25'),
(22, 900.00, '2025-01-25'),
(22, 900.00, '2025-02-25'),
(22, 900.00, '2025-03-25'),
(23, 850.00, '2025-01-25'),
(23, 850.00, '2025-02-25'),
(23, 850.00, '2025-03-25'),
(24, 950.00, '2025-01-25'),
(24, 950.00, '2025-02-25'),
(24, 950.00, '2025-03-25'),
(25, 800.00, '2025-01-25'),
(25, 800.00, '2025-02-25'),
(25, 800.00, '2025-03-25'),
(26, 900.00, '2025-01-25'),
(26, 900.00, '2025-02-25'),
(26, 900.00, '2025-03-25'),

-- Group 5: Vijana Youth
(27, 400.00, '2025-01-05'),
(27, 400.00, '2025-02-05'),
(27, 400.00, '2025-03-05'),
(28, 400.00, '2025-01-05'),
(28, 400.00, '2025-02-05'),
(28, 400.00, '2025-03-05'),
(29, 300.00, '2025-01-05'),
(29, 300.00, '2025-02-05'),
(29, 300.00, '2025-03-05'),
(30, 350.00, '2025-01-05'),
(30, 350.00, '2025-02-05'),
(30, 350.00, '2025-03-05'),
(31, 300.00, '2025-01-05'),
(31, 300.00, '2025-02-05'),
(31, 300.00, '2025-03-05'),
(32, 350.00, '2025-01-05'),
(32, 350.00, '2025-02-05'),
(32, 350.00, '2025-03-05');

-- Insert Loans (realistic purposes and amounts)
INSERT INTO loans (member_id, amount, purpose, status, issued_date, due_date, season_id) VALUES 
-- Group 1: Tumaini VSLA
(3, 5000.00, 'School fees for children', 'active', '2025-02-01', '2025-05-01', 1),
(4, 3000.00, 'Medical expenses', 'active', '2025-02-15', '2025-05-15', 1),
(5, 8000.00, 'Farm inputs - seeds and fertilizer', 'active', '2025-03-01', '2025-06-01', 1),

-- Group 2: Upendo Group
(10, 10000.00, 'Business expansion - retail shop', 'active', '2025-01-25', '2025-04-25', 1),
(11, 4000.00, 'Home improvement - roofing', 'active', '2025-02-10', '2025-05-10', 1),

-- Group 3: Amani Savings
(15, 15000.00, 'Purchase of dairy cow', 'active', '2025-01-15', '2025-07-15', 1),
(16, 6000.00, 'School fees - university', 'active', '2025-02-20', '2025-05-20', 1),
(17, 8000.00, 'Medical treatment - surgery', 'active', '2025-03-05', '2025-06-05', 1),

-- Group 4: Umoja Women
(21, 20000.00, 'Purchase of grinding mill', 'active', '2025-01-30', '2025-07-30', 1),
(22, 12000.00, 'Tailoring business expansion', 'active', '2025-02-15', '2025-05-15', 1),
(23, 5000.00, 'School fees - secondary school', 'active', '2025-03-01', '2025-06-01', 1),

-- Group 5: Vijana Youth
(29, 7000.00, 'Motorcycle purchase for boda boda', 'active', '2025-01-10', '2025-07-10', 1),
(30, 3000.00, 'Hair salon equipment', 'active', '2025-02-05', '2025-05-05', 1);

-- Insert Repayments
INSERT INTO repayments (loan_id, amount, payment_date) VALUES 
-- Loan 1 repayments
(1, 1000.00, '2025-03-01'),
(1, 1000.00, '2025-04-01'),

-- Loan 2 repayments
(2, 1000.00, '2025-03-15'),

-- Loan 3 repayments
(3, 2000.00, '2025-04-01'),

-- Loan 4 repayments
(4, 2500.00, '2025-02-25'),
(4, 2500.00, '2025-03-25'),

-- Loan 5 repayments
(5, 1000.00, '2025-03-10'),

-- Loan 6 repayments
(6, 3000.00, '2025-02-15'),
(6, 3000.00, '2025-03-15'),

-- Loan 7 repayments
(7, 1500.00, '2025-03-20'),

-- Loan 8 repayments
(8, 2000.00, '2025-04-05'),

-- Loan 9 repayments
(9, 5000.00, '2025-02-28'),
(9, 5000.00, '2025-03-28'),

-- Loan 10 repayments
(10, 3000.00, '2025-03-15'),

-- Loan 11 repayments
(11, 1500.00, '2025-04-01'),

-- Loan 12 repayments
(12, 1500.00, '2025-02-10'),
(12, 1500.00, '2025-03-10'),

-- Loan 13 repayments
(13, 1000.00, '2025-03-05');

-- Insert Relief
INSERT INTO relief (member_id, disaster_id, type, amount, date_given) VALUES 
(3, 1, 'Cash Transfer', 2000.00, '2025-07-01'),
(5, 1, 'Cash Transfer', 2000.00, '2025-07-01'),
(10, 2, 'Cash Transfer', 3000.00, '2025-10-01'),
(15, 3, 'Cash Transfer', 2500.00, '2025-08-01'),
(21, 1, 'Cash Transfer', 2000.00, '2025-07-01'),
(29, 1, 'Cash Transfer', 2000.00, '2025-07-01');
