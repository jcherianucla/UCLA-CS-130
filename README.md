# UCLA-CS-130
This houses the entirety of the project for CS 130 - Software Engineering at UCLA, Fall 2017. The site can be found live [here](http://grade-portal.herokuapp.com).

## GradePortal

The idea behind GradePortal is to give students the chance to get real time feedback on projects for a variety of classes as provided by professors.

### The Problem

In many introductory Computer Science classes, students start out with code for the very first time, trying to understand the working environment along with fundamental concepts in Computer Science. They embark on many projects, only to receive a number that tells them how well they did in the project. This number is not enough for constructive feedback and learning as it doesn't let the student know where they failed/succeeded, and what they could do to improve.

### The Solution

Professors in said classes usually have a grading script with test cases that describe which areas the students failed/succeeded in. However, these scripts are usually convoluted to the point that if a student tries to run it themselves on their project - post submission - they aren't going to be able to digest the information effectively. GradePortal is the solution to this very problem. It provides a two sided interface, for professors to easily create projects, administer grades, upload grading scripts etc. and thereby allows a student to upload their submission to the portal to get immediate feedback on the results of the grading script.

### Scope

For the purpose of CS 130, the scope of this project will primarily focus on creating a robust platform for CS 31 and CS 32.

## Technology Stack

GradePortal will be a web application (to make it accessible to any individual) that uses Bruin Logon (through Google OmniAuth) to ensure access to UCLA students. To make the portal extensible for future work, the architecture employed will be a SOA with the use of traditional MVC (Model-View-Controller) functionality exposed through a RESTful API.

* Frontend - React
* Backend - Golang

## Documentation

* [Frontend](https://github.com/jcherianucla/UCLA-CS-130/blob/master/frontend/README.md)
* [Backend](https://github.com/jcherianucla/UCLA-CS-130/blob/master/api/README.md)

## The Team

* [Katie Aspinwall](https://github.com/kaspii)
* [Shalini Dangi](https://github.com/shalinidangi)
* [Connor Kenny](https://github.com/ckenny9739)
* [Jahan Kuruvilla Cherian](https://github.com/jcherianucla)
* [Omar Ozgur](https://github.com/omar-ozgur)
