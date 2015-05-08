# Overriding-Server

This project is extremely simple, and implements a golang web server that serves files from two directories, giving priority to one. This is intended to allow easy pulls from the origin repository of a library while overlaying changes. For example, a project directory would look like the following:

- My Project
  - server.exe
  - base
    - blockly (unmodified from gitHub)
    	- core
    	- examples
    - closure-library (unmodified from gitHub)
  - overrides
    - blockly
    	- core
    		- [a couple files]
    	- examples
    		- [My new example]
    		
In this case, a request to http://localhost:8000/blockly/core/[something.html] will return from the overrides directory if that file exists there, otherwise it will return from the base directory.

At the moment, no customization of the serving port or directory names is possible (sorry), so watch out for conflicts with Skype.
