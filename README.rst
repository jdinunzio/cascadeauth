cascadeauth
===========

cascadeauth is a squid authenticator. It is used to authenticate to
multiple sources, using others authenticators. If any of them answers "OK",
then it answers "OK".


Invocation
==========

Invoke cascadeauth with the path of its config file::

    $ cascadeauth /path/to/config.file


Requirements
============

cascadeauth requires the go programming language and gnu make.


Compilation
===========

Just type::

    make clean; make


Instalation
===========

Copy cascadeauth to your /usr/bin directory


Configuration
=============

The configuration file consist in a serie of lines, each one is an invocation 
to a squid authenticator::

    /path/to/authenticator1
    /path/to/authenticator2

At the moment, it doesn't accept comments or blank lines

