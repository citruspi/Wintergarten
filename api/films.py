import falcon
import json
import os
import requests
from . import config

try:

    caching = config.getboolean('Redis', 'cache')

    if caching:

        import redis

        try:

            cache = redis.StrictRedis(host=config.get('Redis', 'host'),
                                      port=config.get('Redis', 'port'),
                                      db=config.get('Redis', 'db'),
                                      password=config.get('Redis', 'password'))

        except Exception:

            cache = redis.StrictRedis(host=config.get('Redis', 'host'),
                                      port=config.get('Redis', 'port'),
                                      db=config.get('Redis', 'db'))

except Exception, e:

    print e
    caching = False

class FilmItem(object):

    def on_get (self, req, resp, id):

        if caching:

            key = config.get('Redis', 'namespace') + '_film_title_'+id

            if cache.exists(key):

                film = cache.get(key)

                resp.status = falcon.HTTP_200
                resp.body = film

                return

        TMDB_API_KEY = config.get('TheMovieDB', 'API_KEY')

        extra = 'credits,images,releases,similar_movies,reviews'

        r = requests.get('http://api.themoviedb.org/3/movie/'+id,
                            params={
                                'api_key': TMDB_API_KEY,
                                'append_to_response': extra
                            })

        if r.status_code == 200:

            film = r.json()

            film = {}

            film['info'] = r.json()

            r = requests.get('http://www.canistream.it/services/search',
                              params={
                                'movieName': film['info']['title']
                              },
                              headers={
                                'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.43.40 (KHTML, like Gecko) Version/8.0 Safari/538.43.40'
                              })

            if r.status_code == 200:

                films = r.json()

                for f in films:

                    if 'imdb' in f['links']:

                        imdb_id = f['links']['imdb'].split('/')[-2]

                        if film['info']['imdb_id'] == imdb_id:

                            film['availability'] = {}

                            for media in ['streaming', 'rental', 'purchase', 'dvd', 'xfinity']:

                                r = requests.get('http://www.canistream.it/services/query',
                                                    params={
                                                        'movieId': f['_id'],
                                                        'attributes': 1,
                                                        'mediaType': media
                                                    },
                                                    headers={
                                                        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10) AppleWebKit/538.43.40 (KHTML, like Gecko) Version/8.0 Safari/538.43.40'
                                                    })

                                if r.status_code == 200:

                                    film['availability'][media] = r.json()

            if caching:

                key = config.get('Redis', 'namespace') + '_film_title_'+id

                cache.set(key, json.dumps(film))
                cache.expire(key, config.get('Redis', 'lifetime'))

            resp.status = falcon.HTTP_200
            resp.body = json.dumps(film)

        elif r.status_code == 404:

            resp.status = falcon.HTTP_404
            resp.body = ''

        else:

            resp.status = falcon.HTTP_500
            resp.body = ''

class FilmSearch(object):

    def on_get (self, req, resp, query, page=1):

        TMDB_API_KEY = config.get('TheMovieDB', 'API_KEY')

        r = requests.get('http://api.themoviedb.org/3/search/movie',
                            params={
                                'api_key': TMDB_API_KEY,
                                'query': query,
                                'page': page
                            })

        if r.status_code == 200:

            result = r.json()

            resp.status = falcon.HTTP_200
            resp.body = json.dumps(result)

        elif r.status_code == 404:

            resp.status = falcon.HTTP_404
            resp.body = ''

        else:

            resp.status = falcon.HTTP_500
            resp.body = ''

class FilmSet (object):

    def on_get (self, req, resp, set, page=1):

        sets = ['latest', 'upcoming', 'now_playing', 'top_rated', 'popular']

        if set not in sets:

            resp.status = falcon.HTTP_404
            resp.body = ''

            return

        TMDB_API_KEY = config.get('TheMovieDB', 'API_KEY')

        r = requests.get('http://api.themoviedb.org/3/movie/' + set, params={
            'api_key': TMDB_API_KEY,
            'page': page
        })

        if r.status_code == 200:

            result = r.json()

            resp.status = falcon.HTTP_200
            resp.body = json.dumps(result)

        elif r.status_code == 404:

            resp.status = falcon.HTTP_404
            resp.body = ''

        else:

            resp.status = falcon.HTTP_500
            resp.body = ''
