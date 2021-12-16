# ScriFlow
A web-based workflow system for [Scribus](https://en.wikipedia.org/wiki/Scribus)*

** This software is very, very alpha version-ish - expect nothing to work at the moment, download the source if you want to hack on it. **

This project is an attempt to make a workflow system, allowing multiple persons to work simultaneously on a Scribus project.
I stumbled across Scribus many moons ago. However it wasn't until recently that I had the chance to actually use it for a real project.   
After that time I thought to myself that it would be cool to try to build an ecosystem around it, just as people have done with commercially available products. But whereas traditional companies often try to force vendor lockins upon its users, we do the opposite: Let people maintain their software freedom!

*) Not affiliated with the main Scribus project. I am, however, a happy Scribus user. (^Simon) 

## Project Outline
The project is still in its infancy, and consists of two parts, a React-powered clientside app being the frontend GUI, and a lightweight server written in Go. The server will act as "middleware" maintaining connections for authenticated project users while providing realtime updates via the WebSocket protocol

## Contributions always welcome
You're always welcome to participate in the project. Download the code and start hacking. In order to avoid that people are working on the same features,
it might be a good idea to post your ideas in the discussion forum before getting knee-deep in work.

## Endgoals
* Create a CRUD system to manage projects
* Support the same project types as Scribus
* Make the app look and feel as close as possible a native offline app.
* Create a permission system allowing for finegrained control regarding what actions project participants is allowed to make.
* Create a media library so users can share assets between projects
* Allow authenticated users to control the canvas	
* Role-based authentication (based on go module Authority)
* Extend Scribus with a plugin to support ScriFlow (this would be huge!]
* Extend Scribus and ScriFlow to be able to use cloud storage platforms like NextCloud
