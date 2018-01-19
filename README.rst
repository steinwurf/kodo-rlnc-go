kodo-rlnc-go
============

Go bindings for `kodo-rlnc-c`_.

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

To check if your installation was success you can try to build the
``encode_decode_example`` and run it like so:

::

    $GOPATH/bin/encode_decode_example

.. Running the tests
    -----------------

    Explain how to run the automated tests for this system

    Break down into end to end tests
    ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

    Explain what these tests test and why

    ::

        Give an example

    And coding style tests
    ~~~~~~~~~~~~~~~~~~~~~~

    Explain what these tests test and why

    ::

        Give an example

    Deployment
    ----------

    Add additional notes about how to deploy this on a live system

    Built With
    ----------

    -  `Dropwizard`_ - The web framework used
    -  `Maven`_ - Dependency Management
    -  `ROME`_ - Used to generate RSS Feeds

    Contributing
    ------------

    Please read `CONTRIBUTING.md`_ for details on our code of conduct, and
    the process for submitting pull requests to us.

    Versioning
    ----------

    We use `SemVer`_ for versioning. For the versions available, see the
    `tags on this repository`_.

    Authors
    -------

    -  **Billie Thompson** - *Initial work* - `PurpleBooth`_

    See also the list of `contributors`_ who participated in this project.

    License
    -------

    This project is licensed under the MIT License - see the `LICENSE.md`_
    file for details

    Acknowledgments
    ---------------

    -  Hat tip to anyone who’s code was used
    -  Inspiration
    -  etc

.. _kodo-rlnc-c: https://github.com/steinwurf/kodo-rlnc-c
.. _Dropwizard: http://www.dropwizard.io/1.0.2/docs/
.. _Maven: https://maven.apache.org/
.. _ROME: https://rometools.github.io/rome/
.. _CONTRIBUTING.md: https://gist.github.com/PurpleBooth/b24679402957c63ec426
.. _SemVer: http://semver.org/
.. _tags on this repository: https://github.com/your/project/tags
.. _PurpleBooth: https://github.com/PurpleBooth
.. _contributors: https://github.com/your/project/contributors
.. _LICENSE.md: LICENSE.md