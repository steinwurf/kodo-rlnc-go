kodo-rlnc-go
============

Go bindings for `kodo-rlnc-c`_.

.. _kodo-rlnc-c: https://github.com/steinwurf/kodo-rlnc-c

Getting Started
---------------

These instructions will get you ready to start using this project
as a dependency for your go projects.

Prerequisites
~~~~~~~~~~~~~

Since these bindings depends on a C library which is not available from
a package, you need to build and make resulting binaries available
before you can ``go get`` this library.

First checkout this git project:
::

    git clone https://github.com/steinwurf/kodo-rlnc-go


Configure and build the project.
::

    cd kodo-rlnc-go
    python waf configure
    python waf build

After a success compilation the products of the build needs to be made
available to kodo-rlnc-go. This is accomplished with the following waf command.
Make sure you have set your $GOPATH environment variable.
::

    python waf --install_static_libs --install_path $GOPATH/src/github.com/steinwurf/kodo-rlnc-c

``$GOPATH/src/github.com/steinwurf/kodo-rlnc-c`` is the path were kodo-rlnc-go,
when fetched with ``go get``, expects the needed c library and header is
located.

Installing
~~~~~~~~~~

After completing the steps specified in the Prerequisites section, installing
kodo-rlnc-go is as simple using the following ``go get`` command:

::

    go get github.com/steinwurf/kodo-rlnc-go

And similarly it can be used as a dependency like so:

::

    import (
        ...
        "github.com/steinwurf/kodo-rlnc-go"
    )

When using kodo-rlnc-go as a dependency in your project, the
Prerequisites has to be fulfilled before your project can be built.

Running the tests
-----------------

To check if your installation was success you can try to run the tests like so:

::

    Give an example

License
-------

To obtain a valid Kodo license **you must fill out the license request** form_.

Kodo is available under a research- and education-friendly license, see the
details in the LICENSE.rst file.

.. _form: http://steinwurf.com/license/
