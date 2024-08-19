# Flask
Flask is a web framework used for Python which offers simplicity and ease of use.

## Production Deployment
Some examples of tools used to deploy Flask in production are
- Gunicorm
- uWSGI
- Gevent
- Twsited Web

To run an application named as main.py with the flask app component, named as app, run the below command;

`gunicorn main:app`

The default port is 8000. You can also add multiple workers by adding the -w flag with the desired number of workers. An example can be seen below;

`gunicorn main:web -w 2`