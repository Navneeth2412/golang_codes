You need to develop a backend for the Student Marks Processing application.

Student marks are shared in StudentMarks.CSV, write a “Bulk Load” component in a separate
file to read the content of CSV and insert into the MySQL database Student Marks table.

Write another component called “Marks Processor” to process marks data present in the
MySQL database. If a student is getting more than or equal to 70 marks, the pass result should
be “Pass with Distinction”, and less than 70 and more than 40 should be “Pass”. Create and
insert into another table with results data.

Write another component called “Web Service” to expose Restful API for get request. If
student id is passed in get request, result should be sent back.
