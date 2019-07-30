#! /usr/bin/env python
# encoding: utf-8
import os
import string

from waflib import Utils
from waflib.TaskGen import feature, after_method

APPNAME = 'kodo-rlnc-go'
VERSION = '1.0.0'


def build(bld):
    if not bld.has_tool_option('install_static_libs'):
        return

    # We assume that install_path is available if install_static_libs is set.
    install_path = bld.get_tool_option('install_path')
    install_path = os.path.abspath(os.path.expanduser(install_path))

    # Substitute environment variables
    install_path = string.Template(install_path).substitute(os.environ)
    kodo_rlnc_c = os.path.realpath(bld.dependency_path('kodo-rlnc-c'))

    headers = ['encoder.h', 'decoder.h', 'common.h']
    install_files = []
    for header in headers:
        path = os.path.join(kodo_rlnc_c, 'src', 'kodo_rlnc_c', header)
        install_files.append(bld.root.find_node(path))

    bld.install_files(install_path, install_files)
