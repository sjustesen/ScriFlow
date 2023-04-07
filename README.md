# ScriFlow (working title)
A DTP workflow system for the web

** This software is very, very alpha version-ish - expect nothing to work at the moment, download the source if you want to hack on it. 
**
This project is an attempt to make a workflow system, allowing multiple persons to work simultaneously on a DTP project.

## Project Outline
The consists of two parts, a React-powered clientside app being the frontend GUI, and a lightweight server written in Go. 
The server will act as a backend maintaining connections for authenticated project users while also serving REST requests and responses.
We will also explore the opportunity to be providing realtime updates back and fourth via the WebSocket protocol

## Contributions always welcome
You're always welcome to participate in the project. Download the code and start hacking. In order to avoid that people are working on the same features,
it might be a good idea to post your ideas in the discussion forum before getting knee-deep in work.

## Goals
* Create a CRUD system to manage projects
* Support the same filetypes as open systems like [Scribus](https://en.wikipedia.org/wiki/Scribus), Collabora Office, etc.
* Make the app look and feel as close as possible a native offline app.
* Allow authenticated users to control the canvas	
* Extend [Scribus](https://en.wikipedia.org/wiki/Scribus) with a plugin to support uploading/downloading to/from ScriFlow
* Enable ScriFlow to be able to use cloud platforms like NextCloud for storage
