kodo-rlnc-go
============

Go bindings for `kodo-rlnc-c`_.


Getting Started
---------------

These instructions will get you ready to start using this project
as a dependency for your go projects.

Prerequisites
~~~~~~~~~~~~~

This project depends on `kodo-rlnc-c`_, a C library which is not available as
a package. For this reason you will need to build and install `kodo-rlnc-c`_
before a successful executing of ``go get github.com/steinwurf/kodo-rlnc-go``
can be performed.

First checkout this git project.

::

    git clone https://github.com/steinwurf/kodo-rlnc-go


Use Waf to configure and build. This will ensure the correct version of
`kodo-rlnc-c`_ is used.

::

    cd kodo-rlnc-go
    python waf configure
    python waf build

After a successful configuration and compilation the products of the build
needs to be made available. This is accomplished with the following Waf install
command. Make sure you have set your $GOPATH environment variable.

::

    python waf install --install_static_libs --install_path $GOPATH/src/github.com/steinwurf/kodo-rlnc-c

``$GOPATH/src/github.com/steinwurf/kodo-rlnc-c`` is the path were kodo-rlnc-go
expects the needed static library and header is located.

Installing
~~~~~~~~~~

After completing the steps specified in `Prerequisites`_, installing
kodo-rlnc-go is as simple using the following ``go get`` command:

::

    go get github.com/steinwurf/kodo-rlnc-go

And similarly it can be used as a dependency like so:

::

    import (
        ...
        "github.com/steinwurf/kodo-rlnc-go"
    )

When using kodo-rlnc-go as a dependency in your project, the directions in
`Prerequisites`_ needs to be followed before your project can be built.

Development
-----------

If you want to extend the bindings with new features or fix an issue,
follow the steps specified in Prerequisites and Installing.
When you have completed these steps you can and make, test and commit your
changes from the ``$GOPATH/src/github.com/steinwurf/kodo-rlnc-go`` directory.

Running the tests
-----------------

To check if your installation was successful you can try to run the tests like
so:

::
    cd $GOPATH/src/github.com/steinwurf/kodo-rlnc-go
    go test

License
-------
You will need a valid license to build `kodo-rlnc-c`_.

To obtain a valid Kodo license **you must fill out the license request** form_.

Kodo is available under a research- and education-friendly license, see the
details in the LICENSE.rst file.

.. _form: http://steinwurf.com/license/
.. _kodo-rlnc-c: https://github.com/steinwurf/kodo-rlnc-c
