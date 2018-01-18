#! /usr/bin/env python
# encoding: utf-8
import sys
import os
import string

from waflib import Utils
from waflib.TaskGen import feature, after_method

APPNAME = 'kodo-rlnc-go'
VERSION = '0.0.0'

# def run_cmd(cwd, cmd):
#     #todo run these commands after building libraries
#     print("cd " + cwd)
#     print(" ".join(cmd))


def options(opt):
    opts = opt.get_option_group('Install path options')
    opts.add_option(
        '--kodo_rlnc_c_header_install_path', default=None, dest='kodo_rlnc_c_header_install_path',
        help="Install path for the kodo_rlnc_c.h header file.")


def configure(conf):
    conf.find_program('go', var='GO')


def build(bld):


    if bld.has_tool_option('install_path') and \
       bld.has_tool_option('install_static_libs'):
        install_path = bld.get_tool_option('install_path')
        install_path = os.path.abspath(os.path.expanduser(install_path))
        install_path = string.Template(install_path).substitute(os.environ)

        kodo_rlnc_src = os.path.realpath(os.path.join(bld.dependency_path('kodo-rlnc-c'), 'src', 'kodo_rlnc_c'))
        kodo_rlnc_src = bld.root.find_dir(kodo_rlnc_src)
        kodo_rlnc_c_h = kodo_rlnc_src.find_node('kodo_rlnc_c.h')

        bld.install_files(install_path, [kodo_rlnc_c_h])

    # populate these two variables with the correct information from bld
    # includes = "/path/to/kodo_rlnc_c.h"
    # static_libs = {
    #     "/static/lib/path1": "my_static_lib1",
    #     "/static/lib/path2": "my_static_lib2",
    #     "/static/lib/path3": "my_static_lib3"
    # }

    # cgo_cflags = "CGO_CFLAGS=\"-I{includes}\"".format(includes=includes)

    # static_libs_string = ''.join("-L{} -l{} ".format(p,l) for p, l in static_libs.items()).strip()
    # cgo_ldflags = "CGO_LDFLAGS=\"{static_libs}\"".format(static_libs=static_libs_string)

    # go = bld.env.get_flat('GO')
    # cmd = [cgo_cflags, cgo_ldflags, go, 'build']

    # run_cmd(cwd="kodorlnc/", cmd=cmd)
