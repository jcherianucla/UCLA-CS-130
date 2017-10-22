# UCLA-CS-130
This houses the entirety of the project for CS 130 - Software Engineering at UCLA, Fall 2017.

## GradePortal

The idea behind GradePortal is to give students the chance to get real time feedback on projects for a variety of classes as provided by professors.

### The Problem

In many introductory Computer Science classes, students start out with code for the very first time, trying to understand the working environment along with fundamental concepts in Computer Science. They embark on many projects, only to receive a number that tells them how well they did in the project. This number is not enough for constructive feedback and learning as it doesn't let the student know where they failed/succeeded, and what they could do to improve.

### The Solution

Professors in said classes, usually have a grading script with test cases that describe which areas the students failed/succeeded. However, these scripts are usually convoluted to the point that if a student tries to run it themselves on their project - post submission - they aren't going to be able to digest the information effectively. GradePortal is the solution to this very problem. It provied a two sided interface, for professors to easily create projects, administer grades, upload grading scripts etc. and thereby allows a student to upload their submission to the portal to get immediate feedback on the results of the grading script.

### Scope

For the purpose of CS 130, the scope of this project will primarily focus on creating a robust platform for CS 31 and CS 32.

## Technology Stack

GradePortal will be a web application (to make it accessible to any individual) that uses Bruin Logon (through Google OmniAuth) to ensure access to UCLA students. To make the portal extensible for future work, the architecture employed will be a SOA with the use of traditional MVC (Model-View-Controller) functionality exposed through a RESTful API.

* Frontend - React
* Backend - Golang

## The Team

* Katie Aspinwall
* Shalini Dangi
* Prit Joshi
* Connor Kenny
* Jahan Kuruvilla Cherian
* Omar Ozgur
