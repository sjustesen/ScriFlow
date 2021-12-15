# WebWorkflowWithScribus
Very much work in progress - expect everything to be experimental and nonfunctioning ;) Merry XMas!

Scribus - an online DTP app for various desktops - can act as a server.
You feed into it a template on the command line and then it outputs e.g. ä PDF file.

This project is an attempt to make a workflow system, allowing multiple persons to work on a project simultaneously

## Project Outline

The project consists of two parts, a React-powered clientside app being the frontend GUI, and a lightweight server written in Go. The server will act as "middleware" maintaining connections for authenticated project users while providing realtime updates via the WebSocket protocol

## Endgoals
* Create a CRUD system to manage projects
* Support the same project types as Scribus
* Make the app look and feel as close as possible a native offline app.
* Create a permission system allowing for finegrained control regarding what actions a project participant is allowed to make.
* Create a media library so users can share assets between projects
* Allow authenticated users to control the canvas	
* Role-based authentication (based on go module Authority)

## Features
to be decided
