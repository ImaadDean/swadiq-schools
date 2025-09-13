#!/bin/bash

echo "Student API Testing Script"
echo "=========================="
echo ""

BASE_URL="http://localhost:3000/api/students"

echo "1. Get all students (table format):"
echo "GET $BASE_URL/table"
echo ""

echo "2. Get students by year 2025:"
echo "GET $BASE_URL/year?year=2025"
echo ""

echo "3. Get all students (basic format):"
echo "GET $BASE_URL"
echo ""

echo "4. Get students by class:"
echo "GET $BASE_URL/class?class_id=YOUR_CLASS_ID"
echo ""

echo "5. Get single student by ID:"
echo "GET $BASE_URL/STUDENT_ID"
echo ""

echo "6. Create new student:"
echo "POST $BASE_URL"
echo "Content-Type: application/json"
echo ""
echo "Example JSON body:"
cat << 'EOF'
{
  "first_name": "John",
  "last_name": "Doe",
  "date_of_birth": "2010-05-15",
  "gender": "male",
  "address": "123 Main St, City",
  "class_id": "class-uuid-here",
  "parent_id": "parent-uuid-here"
}
EOF

echo ""
echo ""
echo "Note: You need to be authenticated to access these endpoints."
echo "The student_id will be auto-generated in format STU-YYYY-XXX"
