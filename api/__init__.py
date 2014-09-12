import ConfigParser
import os
import sys

config = ConfigParser.RawConfigParser()

if os.path.isfile('wintergarten.conf'):

    config.read('wintergarten.conf')

elif os.path.isfile('/etc/wintergarten.conf'):

    config.read('/etc/wintergarten.conf')

else:

    print 'No configuration file available.'
    sys.exit(1)

from .films import FilmItem, FilmSearch, FilmSet
