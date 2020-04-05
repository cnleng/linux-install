heroku login -i

heroku create myapp --buildpack heroku/python

heroku addons:create heroku-postgresql:hobby-dev

heroku buildpacks:set heroku/python

heroku apps:destroy
